package snark

import (
	"encoding/hex"
	"log"
	"testing"
)

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}

func TestVerify(t *testing.T) {
	proof := "8519e84088a58ab794f4b50ae7b3790e880421d7fc129ceeb2314059a343d5d3d20e4cd3c03769c700748ae0bae0c5a1e937ef2273eda521ce9112a20c3c6d9f82b455aa6e74b65b9e8b32971c2eba733b96c2f33cef65caa196018679e94a002ee006f572625152c8610bd9ee02f7c3086353a3c19ba0039205b4b170a0b5f1026b3c108af4dd0c329a1b693b0f4ed8e1820a1b9ffb935ada69e5c503e5634df3f63b49be617fd260bbaa1518c4acbc0a3c6663275fca1faf8103342ee798002f974b72480b9f3cad18b9e8a8371c34d465e119f02aec8fd5c7c9eed108eb7b516685fa0781fd929affb79a469a5fc1889998fdf252dab2230134a0f2b0171547ef81381c1d5b719e72e1341afe11f09a752b700c7f5308501dcc4f451ca080"
	vk := "8a02409a340bc61af1dabbb2e8ea92505d65e572889218e4b23ee4e97f95dee3e35a53ecb08b939f22fa5b2ee0e7ce965d7d4a482acde57c7347859b5f997f1cd28b47c5d6b57e567b8806e7cc5b04950f1cbabcc0c8efedf6a67f2bd5fca8004707d7f267382e490928922f4f7a23f4cf245134198c6c581a253631cb95bb6ebe5754ee79fcfcb3182b42e98547f9c49b05d109f3b948f21c59c19e32ddb3fc2f32976995413715ccb153e1b177c575eed8a91d7d4fd63c710b7fe9a34222803dc77891311f2fa5d31043174a18b1619b5852a960fafa3f4bb6a7e29aa8ab394cd67ab8ced0beb297d4c1ace7af932296b438be3a02fa5689a45c7697159f58398be049b9e7b05d72f10d3ad4688c755021a5ea096e6a57bd37015b28ad2500f6a0607b8a566e73629db5be2ba902f1640b9100d8fab6dfd5e7cd90f62462609efd610bc65fa2418781697a7a8d93cb6984073d08efdeeeffb9884f72401feb4ef9d059ba5ccda8da6ae541f41d1a16212feeaf8c618afadfc1c53f86bc28000300000000000000e9119ca55b2d31f037cba7a6f22ca2774054d7e21802f44bd836bd34d8d4a76ff6c44c6c95071ffc4cebc063de8bd13adc86c85af61f965da385f8c84e39eac0a98801c370055103127cad716b3db4da723745044a1e08741e386871a5edb400583d48f6755076c514950e774b5f22d802f5689c45837eef1fdfd26025fee18bac6b2b2b4fb4a564193d65bec77f1233244acd345cd6357df2372a9b093f5407526b92b973b61cdcb91dcf11c0992de37dcf80dde0d1c2a943da0b4d8aed378028782729577756708a3585132eb7fcaef346aaa443a4b303d86e992f6f5bdfb01869480dab8699beb153fd0d2b73cbcdb1e65f98e812bf0640c0fb152c3546000a0a65eda8ff3cf0aef0425b127e7dccc7c3c40fe44970e57faf234e77a4a900"
	first_epoch_pubkeys := "4fd0ba041e8118dd4fe88d7635e04bbd90023b93e468c19bb559c1f91c886b190d86f924777b160184a6f0702e4800013e7932ef27d0039705d0dfccb3af08852acc1ac0ef0a5eae185608ea85fd2dd79eab385f973f150d03d8ccca7a1e0b00034788f77e75a0d11ac446be49c8301126fc31b525f813549e8fb3a0d0f44f9bc75e42e6f0571ab77b62980e831a88019d1eb42131b80cf3d6bab01f284688096bdbfbf340365126127688758837fc06b039de7931430b915b6fe541ad547280040844901a0cad509aa7776050cfd44ac5a4564f7a7d39a955a919e0cb19ddf5b24dc9f2de0c5a2c858c065bf8ad9e00b3c71d9a81b8de74c4dd2e82b535d47edcf9a2d44f1fd1804e194fc1d9b57537f91c09b1cd9d78c3c0c9cc4fd61b8380765e6251d4024a69e0d8d3879fcfd7db00578bbc21371bd1cbe46e96d67b1e2a6178716a0c7383f55ad0eff94e7c3c0010260024ed5b0e76054d85e8add6e8298c05523f7af4669b0b4ba94bba72a45246662cfe1dbb96e7ffe2caac74d57e80"
	last_epoch_pubkeys := "497d9ddf216691884d5e7a1e91d019d4022134f3d402d44db65ae2e607e5a2237a19082953cab8376f8a59b6ccaf6f01356337be9ccb34cceb861667c9eadebb502bf55a4eda682c19df28e215e3cdffab7294c5a321aacb74866203d7870b81fddab661d3e74f31b417596b9ad2b5c2866ec6d2a1c052164cb6fcea7bc49ea3782111b42e1fdf5a26138df9bfbd1200e48988a735f87f4d1e855ae9ce841643d8d5c0f43a2795a250ce1abc87b85e9305231beb2bfd7647766110331deb7401cde82d55537417981333bdb0a9ad49e1d706f1acc6f749e24d71df7e7025bfe58d3d912b1b45c6ab9a7137cbf0aafc00b700814034590a464dacdb57c6f4b7753f455e51d04df5844083e4202a275547a0f70f9bf432bea06c74c1f8f7363301ba713be2e78be900d01e38b8ccae5fef407e9ed5661962e77b96743a6a90d42b2cfa0b7c4d53256069cc3cd969032d00e5c086447248cdec7e5b71e3a248d6162c7f159cb370493787c4d2de8f6135c386c814c07728c4253159b41fe0527580"

	epoch1EpochEntropy := make([]byte, 16)
	epoch1ParentEntropy := make([]byte, 16)
	for i := 0; i < 16; i++ {
		epoch1EpochEntropy[i] = 1
		epoch1ParentEntropy[i] = 2
	}

	pk1, _ := hex.DecodeString(first_epoch_pubkeys)
	publicKeys := split(pk1, PUBLIC_KEY_BYTES)
	epoch1 := EpochBlock{
		Index:         0,
		Round:         0,
		MaxNonSigners: 1,
		MaxValidators: uint(len(publicKeys)),
		EpochEntropy:  epoch1EpochEntropy,
		ParentEntropy: epoch1ParentEntropy,
		PublicKeys:    publicKeys,
	}

	epoch2EpochEntropy := make([]byte, 16)
	epoch2ParentEntropy := make([]byte, 16)
	for i := 0; i < 16; i++ {
		epoch2EpochEntropy[i] = 3
		epoch2ParentEntropy[i] = 2
	}

	pk2, _ := hex.DecodeString(last_epoch_pubkeys)
	epoch2 := EpochBlock{
		Index:         2,
		Round:         0,
		MaxNonSigners: 1,
		MaxValidators: uint(len(publicKeys)),
		EpochEntropy:  epoch2EpochEntropy,
		ParentEntropy: epoch2ParentEntropy,
		PublicKeys:    split(pk2, PUBLIC_KEY_BYTES),
	}

	proofDec, _ := hex.DecodeString(proof)
	vkDec, _ := hex.DecodeString(vk)

	err := VerifyEpochs(
		vkDec,
		proofDec,
		epoch1,
		epoch2,
	)

	if err != nil {
		log.Fatal(err)
	}
}
