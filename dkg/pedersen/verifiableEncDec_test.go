package pedersen

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	_ "net/http/pprof"

	"net/http"

	"go.dedis.ch/dela/dkg"
	"go.dedis.ch/dela/dkg/pedersen/types"
	"go.dedis.ch/dela/mino"

	"go.dedis.ch/dela/mino/minoch"
	_ "go.dedis.ch/dela/mino/minoch"
	"go.dedis.ch/dela/mino/minogrpc"
	"go.dedis.ch/dela/mino/router/tree"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/xof/keccak"
)

func init() {
	rand.Seed(0)
}

func Test_verifiableEncDec_minoch(t *testing.T) {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	// setting up the dkg
	n := 180
	threshold := 180
	batchSize := 100
	workerNum := 8

	minos := make([]mino.Mino, n)
	dkgs := make([]dkg.DKG, n)
	addrs := make([]mino.Address, n)

	// creating GBar. we need a generator in order to follow the encryption and decryption protocol of https://arxiv.org/pdf/2205.08529.pdf /
	// we take an agreed data among the participants and embed it as a point. the result is the generator that we are seeking
	agreedData := make([]byte, 32)
	_, err := rand.Read(agreedData)
	require.NoError(t, err)
	GBar := suite.Point().Embed(agreedData, keccak.New(agreedData))

	minoManager := minoch.NewManager()

	for i := 0; i < n; i++ {
		minoch := minoch.MustCreate(minoManager, fmt.Sprintf("addr %d", i))
		minos[i] = minoch
		addrs[i] = minoch.GetAddress()
	}

	pubkeys := make([]kyber.Point, len(minos))

	for i, mino := range minos {
		dkg, pubkey := NewPedersen(mino)
		dkgs[i] = dkg
		pubkeys[i] = pubkey
	}

	fakeAuthority := NewAuthority(addrs, pubkeys)

	actors := make([]dkg.Actor, n)
	for i := 0; i < n; i++ {
		actor, err := dkgs[i].Listen()
		require.NoError(t, err)
		actors[i] = actor
	}

	fmt.Println("setting up the dkg ...")

	_, err = actors[0].Setup(fakeAuthority, threshold)
	require.NoError(t, err)

	fmt.Println("generating the message and encrypting it ...")
	//generating random messages in batch and encrypt them
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	keys := make([][]byte, batchSize)
	var ciphertexts []types.Ciphertext
	for i := 0; i < batchSize; i++ {
		keys[i] = make([]byte, 29)
		for j := range keys[i] {
			keys[i][j] = letterBytes[rand.Intn(len(letterBytes))]
		}
		ciphertext, remainder, err := actors[0].VerifiableEncrypt(keys[i], GBar)
		require.NoError(t, err)
		require.Len(t, remainder, 0)
		ciphertexts = append(ciphertexts, ciphertext)
	}

	// decrypting the batch ciphertext message
	fmt.Println("decrypting the ciphertext ...")
	decrypted, err := actors[0].VerifiableDecrypt(ciphertexts, workerNum)
	require.NoError(t, err)
	for i := 0; i < batchSize; i++ {
		require.Equal(t, keys[i], decrypted[i])
	}

}

func Test_verifiableEncDec_minogrpc(t *testing.T) {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// we want to time the decryption for different batch sizes with different number of nodes
	// numWorkersSlice := []int{16, 16, 32, 64, 64, 64, 64}
	// batchSizeSlice := []int{32, 64, 128, 256, 512, 1024, 2048}
	numWorkersSlice := []int{8}
	batchSizeSlice := []int{32}

	// initiating the log file for writing the delay and throughput data
	f, err := os.OpenFile("../logs/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	require.NoError(t, err)
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	// setting up the dkg
	n := 64
	threshold := 64

	minos := make([]mino.Mino, n)
	dkgs := make([]dkg.DKG, n)
	addrs := make([]mino.Address, n)

	// creating GBar. we need a generator in order to follow the encryption and decryption protocol of https://arxiv.org/pdf/2205.08529.pdf /
	// we take an agreed data among the participants and embed it as a point. the result is the generator that we are seeking
	agreedData := make([]byte, 32)
	_, err = rand.Read(agreedData)
	require.NoError(t, err)
	GBar := suite.Point().Embed(agreedData, keccak.New(agreedData))

	fmt.Println("initiating the dkg nodes ...")
	for i := 0; i < n; i++ {
		addr := minogrpc.ParseAddress("127.0.0.1", 0)

		minogrpc, err := minogrpc.NewMinogrpc(addr, nil, tree.NewRouter(minogrpc.NewAddressFactory()))
		require.NoError(t, err)

		defer minogrpc.GracefulStop()

		minos[i] = minogrpc
		addrs[i] = minogrpc.GetAddress()
	}

	pubkeys := make([]kyber.Point, len(minos))

	for i, mino := range minos {
		for _, m := range minos {
			mino.(*minogrpc.Minogrpc).GetCertificateStore().Store(m.GetAddress(), m.(*minogrpc.Minogrpc).GetCertificateChain())
		}
		dkg, pubkey := NewPedersen(mino.(*minogrpc.Minogrpc))
		dkgs[i] = dkg
		pubkeys[i] = pubkey
	}

	fakeAuthority := NewAuthority(addrs, pubkeys)

	actors := make([]dkg.Actor, n)
	for i := 0; i < n; i++ {
		actor, err := dkgs[i].Listen()
		require.NoError(t, err)
		actors[i] = actor
	}

	fmt.Println("setting up the dkg ...")
	start := time.Now()
	_, err = actors[0].Setup(fakeAuthority, threshold)
	require.NoError(t, err)
	setupTime := time.Since(start)

	//generating random messages in batch and encrypt them
	for i, batchSize := range batchSizeSlice {
		fmt.Printf("=== starting the process with batch size = %d === \n", batchSize)
		workerNum = numWorkersSlice[i]

		const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		keys := make([][]byte, batchSize)
		var ciphertexts []types.Ciphertext
		for i := 0; i < batchSize; i++ {
			keys[i] = make([]byte, 29)
			for j := range keys[i] {
				keys[i][j] = letterBytes[rand.Intn(len(letterBytes))]
			}
			ciphertext, remainder, err := actors[0].VerifiableEncrypt(keys[i], GBar)
			require.NoError(t, err)
			require.Len(t, remainder, 0)
			ciphertexts = append(ciphertexts, ciphertext)
		}
		// decryopting the batch ciphertext message
		fmt.Println("decrypting the batch ...")
		start = time.Now()
		_, err := actors[0].VerifiableDecrypt(ciphertexts, workerNum)
		decryptionTime := time.Since(start)
		require.NoError(t, err)

		// for i := 0; i < batchSize; i++ {
		// 	require.Equal(t, keys[i], decrypted[i])
		// }

		log.Printf("n = %d , batchSize = %d ,  workerNum = %d,decryption time = %v s, throughput =  %v tx/s , dkg setup time = %v s",
			n, batchSize, workerNum, decryptionTime.Seconds(), float32(batchSize)/float32(decryptionTime.Seconds()), float32(setupTime.Seconds()))

	}

}