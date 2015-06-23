# packge morse

简单的 摩斯码工具库。

提供了默认加密字典 `MOEN`.

使用案例

```go

m := Morse{Dict: MOEN}

en := "-- --- .-. ... .  -.-. --- -.. ."
de := "morse code"

fmt.Printf("Morse: \n%s\n", en)
fmt.Println(m.Decode(en))

fmt.Printf("word: \n%s\n", de)
fmt.Println(m.Encode(de))
fmt.Println(m.Encode("GND"))

```

Dict

![Morse_Code.png](Morse_Code.png)
