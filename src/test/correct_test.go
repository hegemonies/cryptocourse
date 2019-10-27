package test

import (
	DF "cryptocrouse/src/go/Diffie-Hellman"
	elgamalcode "cryptocrouse/src/go/ElGamalCode"
	"cryptocrouse/src/go/EuclideanAlgorithm"
	"cryptocrouse/src/go/FastExp"
	"cryptocrouse/src/go/Fingerprints"
	"cryptocrouse/src/go/RSA"
	"cryptocrouse/src/go/ShamirCode"
	"cryptocrouse/src/go/ShanksAlgorithm"
	"cryptocrouse/src/go/VernamCipher"
	"math/big"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

const (
	MaxNumber = 100000000
	MaxCountTest = 1000
)

func TestFastExpCorrectStable(t *testing.T) {
	//t.Parallel()

	var x, y, m int64 = 13, 17, 15

	in := FastExp.FastExp(x, y, m)
	wait := FastExp.NaivePowWithMod(x, y, m)

	if in != wait {
		t.Errorf("Expected %d, got %d", wait, in)
	}
}

func TestFastExpCorrectRand(t *testing.T) {
	//t.Parallel()

	rand.Seed(time.Now().Unix())

	for i := 0; i < MaxCountTest; i++ {
		x, y, m := rand.Int63n(MaxNumber) + 1, rand.Int63n(MaxNumber) + 1, rand.Int63n(MaxNumber) + 1

		in := FastExp.FastExp(x, y, m)
		wait := FastExp.NaivePowWithMod(x, y, m)

		if in != wait {
			t.Errorf("i = %d Expected %d, got %d (x = %d, y = %d, m = %d)", i, wait, in, x, y, m)
		}
	}
}

func TestSmallFastExpCorrectRand(t *testing.T) {
	//t.Parallel()

	rand.Seed(time.Now().Unix())

	for i := 0; i < MaxCountTest; i++ {
		x, y, m := rand.Uint64() % MaxNumber + 1, rand.Uint64() % MaxNumber + 1, rand.Uint64() % MaxNumber + 1

		in := FastExp.SmallFastExp(x, y, m)
		wait := FastExp.NaivePowWithMod(int64(x), int64(y), int64(m))

		if in != uint64(wait) {
			t.Errorf("i = %d Expected %d, got %d (x = %d, y = %d, m = %d)", i, wait, in, x, y, m)
		}
	}
}

func TestEuclideanAlgoRandStable(t *testing.T) {
	//t.Parallel()

	rand.Seed(time.Now().Unix())

	//countGCD := make(map[uint64]uint64)

	for i := 0; i < MaxCountTest; i++ {
		a, b := rand.Int63n(MaxNumber), rand.Int63n(MaxNumber)

		in, _, _ := EuclideanAlgorithm.GCD(a, b)
		wait := big.NewInt(0).GCD(big.NewInt(0), big.NewInt(0), big.NewInt(int64(a)), big.NewInt(int64(b))).Int64()

		//countGCD[in]++

		if int64(in) != wait {
			t.Errorf("i = %d Expected %d, got %d (a = %d, b = %d)", i, wait, in, a, b)
		}
	}

	//fmt.Printf("count GCD:\n%v\n", countGCD)
}

func TestPrimeFunc(t *testing.T) {
	//t.Parallel()

	var testingNumber uint64 = 2
	for ; testingNumber < MaxCountTest; testingNumber++ {
		in := DF.IsPrime(testingNumber)
		wait := big.NewInt(int64(testingNumber)).ProbablyPrime(0)

		if in != wait {
			t.Errorf("Expected %v, got %v (testing number = %d)", wait, in, testingNumber)
		}
	}
}

func TestConnectionUser(t *testing.T) {
	//t.Parallel()

	cryptoSystem := DF.CryptoSystem{}
	alice := "Alice"
	bob := "Bob"

	for i := 0; i < MaxCountTest; i++ {
		cryptoSystem.Init()
		_ = cryptoSystem.AddUser(alice)
		_ = cryptoSystem.AddUser(bob)
		cryptoSystem.ConnectUsers(alice, bob)

		result := cryptoSystem.CheckConnection(alice, bob)

		if result != true {
			t.Errorf("Expected %v, got %v", true, result)
		}
	}
}

func TestConnectionUserRand(t *testing.T) {
	//t.Parallel()

	rand.Seed(time.Now().Unix())

	cryptoSystem := DF.CryptoSystem{}
	cryptoSystem.Init()

	for i := 0; i < MaxCountTest; i++ {
		userA := strconv.Itoa(rand.Int() % MaxNumber)
		userB := strconv.Itoa(rand.Int() % MaxNumber)

		_ = cryptoSystem.AddUser(userA)
		_ = cryptoSystem.AddUser(userB)
		cryptoSystem.ConnectUsers(userA, userB)

		result := cryptoSystem.CheckConnection(userA, userB)

		if result != true {
			t.Errorf("Expected %v, got %v", true, result)
		}
	}
}

func TestShanksAlgoStableMap(t *testing.T) {
	//t.Parallel()

	var a, waitX, p uint64 = 5, 16, 23
	y := FastExp.SmallFastExp(a, waitX, p)

	inX := ShanksAlgorithm.ShanksAlgo(a, p ,y)

	if waitX != inX {
		testY := FastExp.SmallFastExp(a, inX, p)
		if testY != y {
			t.Errorf("Expected %v, got %v (a = %d, p = %d, y = %d,  testY = %d)", waitX, inX, a, p, y, testY)
		}
	}
}

func TestShanksAlgo3RandMap(t *testing.T) {
	rand.Seed(time.Now().Unix())

	//t.Parallel()

	for i := 0; i < MaxCountTest; i++ {
		a, waitX, p := rand.Uint64() % MaxNumber, rand.Uint64() % MaxNumber, rand.Uint64() % MaxNumber

		if a == 0 || a == 1 {
			a = 2
		}
		if waitX == 0 {
			waitX++
		}
		if p == 0 || p == 1 {
			p = 2
		}

		y := FastExp.SmallFastExp(a, waitX, p)

		inX := ShanksAlgorithm.ShanksAlgo(a, p, y)

		if waitX != inX {
			testY := FastExp.SmallFastExp(a, inX, p)
			if testY != y {
				t.Errorf("Expected %v, got %v (a = %d, p = %d, y = %d,  testY = %d)", waitX, inX, a, p, y, testY)
			}
		}
	}
}

func TestShamirCodeSimple(t *testing.T) {
	producerName := "Alice"
	consumerName := "Bob"

	cryptosystem := ShamirCode.CryptoSystem{}
	cryptosystem.Init()
	_ = cryptosystem.AddUser(producerName)
	_ = cryptosystem.AddUser(consumerName)

	data := []uint64{2, 3, 4, 5, 6, 7, 8, 9}

	cryptosystem.SendMessage(producerName, consumerName, data)

	//cryptosystem.PrintUsers()

	if cryptosystem.CheckMessage(producerName, consumerName) == false {
		producer := cryptosystem.Users[producerName]
		consumer := cryptosystem.Users[consumerName]
		t.Errorf("Expected %v, got %v", producer.GetMessage(), consumer.GetMessage())
	}
}

func TestShamirCodeFile(t *testing.T) {
	producerName := "Alice"
	consumerName := "Bob"

	cryptosystem := ShamirCode.CryptoSystem{}
	cryptosystem.Init()
	_ = cryptosystem.AddUser(producerName)
	_ = cryptosystem.AddUser(consumerName)

	filenameTestData := "test_data.png"

	cryptosystem.SendMessageFromFile(producerName, consumerName, filenameTestData)

	//cryptosystem.PrintUsers()

	if cryptosystem.CheckMessage(producerName, consumerName) == false {
		producer := cryptosystem.Users[producerName]
		consumer := cryptosystem.Users[consumerName]
		t.Errorf("Expected %v, got %v", producer.GetMessage(), consumer.GetMessage())
	}
}

func TestElGamalCodeSimpleData(t *testing.T) {
	producerName := "Alice"
	consumerName := "Bob"

	cryptosystem := elgamalcode.CryptoSystem{}
	cryptosystem.Init()
	_ = cryptosystem.AddUser(producerName)
	_ = cryptosystem.AddUser(consumerName)

	data := []uint64{2, 3, 4, 5, 6, 7, 8, 9}

	cryptosystem.SendMessage(producerName, consumerName, data)

	//cryptosystem.PrintUsers()

	if cryptosystem.CheckMessage(producerName, consumerName) == false {
		producer := cryptosystem.Users[producerName]
		consumer := cryptosystem.Users[consumerName]
		t.Errorf("Expected %v, got %v", producer.GetMessage(), consumer.GetMessage())
	}
}

func TestElGamalCodeOnFile(t *testing.T) {
	producerName := "Alice"
	consumerName := "Bob"

	cryptosystem := elgamalcode.CryptoSystem{}
	cryptosystem.Init()
	_ = cryptosystem.AddUser(producerName)
	_ = cryptosystem.AddUser(consumerName)

	filenameTestData := "test_data.png"

	cryptosystem.SendMessageFromFile(producerName, consumerName, filenameTestData)

	//cryptosystem.PrintUsers()

	if cryptosystem.CheckMessage(producerName, consumerName) == false {
		producer := cryptosystem.Users[producerName]
		consumer := cryptosystem.Users[consumerName]
		t.Errorf("Expected %v, got %v", producer.GetMessage(), consumer.GetMessage())
	}
}

func TestVernamCipherSimpleData(t *testing.T) {
	producerName := "Alice"
	consumerName := "Bob"

	cryptosystem := VernamCipher.CryptoSystem{}
	cryptosystem.Init()
	_ = cryptosystem.AddUser(producerName)
	_ = cryptosystem.AddUser(consumerName)

	data := []byte{2, 3, 4, 5, 6, 7, 8, 9}

	cryptosystem.SendMessage(producerName, consumerName, data)

	//cryptosystem.PrintUsers()

	if cryptosystem.CheckMessage(producerName, consumerName) == false {
		producer := cryptosystem.Users[producerName]
		consumer := cryptosystem.Users[consumerName]
		t.Errorf("Expected %v, got %v", producer.GetOriginMessage(), consumer.GetOriginMessage())
	}
}

func TestVernamCipherDataFromFile(t *testing.T) {
	producerName := "Alice"
	consumerName := "Bob"

	cryptosystem := VernamCipher.CryptoSystem{}
	cryptosystem.Init()
	_ = cryptosystem.AddUser(producerName)
	_ = cryptosystem.AddUser(consumerName)

	filenameTestData := "test_data.png"

	cryptosystem.SendMessageFromFile(producerName, consumerName, filenameTestData)

	//cryptosystem.PrintUsers()

	if cryptosystem.CheckMessage(producerName, consumerName) == false {
		producer := cryptosystem.Users[producerName]
		consumer := cryptosystem.Users[consumerName]
		t.Errorf("Expected %v, got %v", producer.GetOriginMessage(), consumer.GetOriginMessage())
	}
}

func TestRSASimpleData(t *testing.T) {
	producerName := "Alice"
	consumerName := "Bob"

	cryptosystem := RSA.CryptoSystem{}
	cryptosystem.Init()
	_ = cryptosystem.AddUser(producerName)
	_ = cryptosystem.AddUser(consumerName)

	data := []uint64{2, 3, 4, 5, 6, 7, 8, 9}

	cryptosystem.SendMessage(producerName, consumerName, data)

	//cryptosystem.PrintUsers()

	if cryptosystem.CheckMessage(producerName, consumerName) == false {
		producer := cryptosystem.Users[producerName]
		consumer := cryptosystem.Users[consumerName]
		t.Errorf("Expected %v, got %v", producer.GetOrigMessage(), consumer.GetOrigMessage())
	}
}

func TestRSADataFromFile(t *testing.T) {
	producerName := "Alice"
	consumerName := "Bob"

	cryptosystem := RSA.CryptoSystem{}
	cryptosystem.Init()
	_ = cryptosystem.AddUser(producerName)
	_ = cryptosystem.AddUser(consumerName)

	filenameTestData := "test_data.png"

	cryptosystem.SendMessageFromFile(producerName, consumerName, filenameTestData)

	//cryptosystem.PrintUsers()

	if cryptosystem.CheckMessage(producerName, consumerName) == false {
		//producer := cryptosystem.Users[producerName]
		//consumer := cryptosystem.Users[consumerName]
		//t.Errorf("Expected %v, got %v", producer.GetOrigMessage(), consumer.GetOrigMessage())
		t.Error("Error")
	}
}

func TestSignatureRSA(t *testing.T) {
	rand.Seed(1)

	userName := "Alice"
	user := Fingerprints.User{}
	user.Name = userName
	user.GeneratePrivateVariables()

	filename := "test_data.png"

	user.ComputeHash(filename)
	user.ComputeSignature()
	//user.WriteHahSumToFile(filename + ".sig")

	//user.PrintUserInfo("name = %s\nc =   %d\nd =   %d\np =   %d\nq =   %d\nn =   %d\nphi =  %d\ny = %v\ns = %v\nhash = %v\n")

	//user.PrintOpenKeysInFile("open-keys-sig-rsa.txt")

	wait := true
	in := user.CheckSignature()

	if wait != in {
		t.Errorf("Expected %v, got %v", wait, in)
	}
}

func TestSignatureElGamal(t *testing.T) {
	//rand.Seed(1)
	t.SkipNow()
	rand.Seed(time.Now().Unix())
	filename := "test_data.png"

	userA := Fingerprints.ElGamalUser{}
	userA.GeneratePrivateNumbers()
	userA.ComputeHash(filename)
	userA.CheckHash()
	userA.ComputeSignature()

	userB := Fingerprints.ElGamalUser{}

	in := userB.CheckMessage(filename, &userA.Y, &userA.R, &userA.S, &userA.G, &userA.P)
	wait := true

	userA.PrintInfo()
	//userB.PrintInfo()

	if wait != in {
		t.Errorf("Expected %v, got %v", wait, in)
	}
}
