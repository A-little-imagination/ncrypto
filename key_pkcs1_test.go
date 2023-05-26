package ncrypto_test

//
//import (
//	"github.com/smartwalle/ncrypto"
//	"testing"
//)
//
//func TestDecodeRSAPKCS1PrivateKey(t *testing.T) {
//	var key = "MIIEpAIBAAKCAQEA+AoH2rRQBgfu4yz7lUIhAjhMoCKi+sMks7QGd4tC/pdh8cUx9c5zhiRDsxP5uQ2b2h1iQ8SyVBX63N8sJcJlIPa2nz1q+hlpx0U8ntdC1Cqy5ytPeqDzvCVQj1kk2dfAOe+k9aPCBTiwuX71tF28bi8vTQkSWAgxIY4IHVgy2Ej5EQcqkLPVht9G/JWMNhZ8t6PzWYnHqTdJDhoh5HklfnTMc5Amdf14ci+cM0Bf6R1uqmnurMxYixhGR+jAJyWPXsuIocMsviirl9HNsllALzGG4YVvygDmQh6ekha3ZBINxsShAVwQjbROAq1HsHMYTRBQclto+L+gkO5EFmhTuwIDAQABAoIBAQDHaMkHuw7aymXznaZlSNTgxAJWe4Irt9c54fGRChTCOVI03SRdSaz2mxESV2zcqWQy5oypjukPUNxt1b3YXYCQD8iiHO50QxoXakJiMQ05JVrPHrK3E87f49/wXuqReDLCYTdthEnniLwLZVaDuxhSPPz6IeDthSh6oYVEkjPOTMc7mLeU8GMbAWp26EaOePK/IFonttsg+CdSpzlgACLie1d2cz9BZVH8HxxjsvZI23wdaMwuO911bLXrNOTaJkv8VtnbXzWPAuPLCMT8HQD+NjqorY226aHm1v0dGM2qZh2kafiZjoKcwJswUq9u0B7YElrbHrxaZqhh+MBjmoYBAoGBAP8jTAZ10oZy6Pocxg4US7f0BXqtwuj/PYRuuiUp9WAs/u4FLJrRD0IJlOinEHhCDN6vprTxkbxv7ppVm+3G3+zZcAhNVPx99aH9okux4R+Kt1uHW04YYiiB/DiOZZjJXiFrmSv+8/F1DCIGI8E8QLKhMRBsTlLw0ihlgHPeR8IxAoGBAPjgl9TSJ2RS1RQVSGRKEGsXBQtALcjbyp3u1T1LxAxcRmYUd/arq/YIRz6IxQpsSL6okR0ql5vJ874Fhvi09IWCxzhA35uKokVRNCwTHgpGw78ke85VihjcO7kFh2X0Za1VxMaynDnwmtbvBUoery6rPovBXAjFSguGJZSwBy2rAoGAVfdMrHiSoN6jBXMggKnyAV9lYTqfC+bEkZTFf7Qti2Rd2VFBAEEkxPl5hRKViA+pSncg0qL6meWVHcdurKVv2dxB1WpJEnJWi2hsgNjEo94Xfe2TpDKvEDsMAvpn3R0sbFJW3+4dv+PSXLx426kzm3yYjZIQqf2esiHRMDlbMGECgYEAqQzFlUX6g+bP3YhnV9gHNva38mBxuRGym3xzU+N3E35KQ2R76A3MDJ6q5gs47E+JShhuGYWo3aVb7yMJ/z3LKjUe5VMLkDem6bay3Y1OxmQy6k752bm4yBIwbwkWY2aS+h5cPytret+4DW7mBOa6z9R586wvu7VnClTgihLXjWkCgYBnYK8a+iMWbYUBRZCxKY4TvENvCo/vakgkQUV5XCdZe7L4ce9/GT6NeWVqiNsxYH9XTgeP04U5aIRPr3rWOk+GcNgSx42XWW4wyJNKRO2BYooOnGPCM5GirtMCf50mpr6YvbEUkB7nqHaVkGMXy9MGZwX3Nzl8e53rff5FmzJXVQ=="
//	if _, err := ncrypto.DecodeRSAPKCS1PrivateKey(ncrypto.FormatPKCS1PrivateKey(key)); err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestDecodeRSAPKCS1PublicKey(t *testing.T) {
//	var key = "MIIBCgKCAQEA+AoH2rRQBgfu4yz7lUIhAjhMoCKi+sMks7QGd4tC/pdh8cUx9c5zhiRDsxP5uQ2b2h1iQ8SyVBX63N8sJcJlIPa2nz1q+hlpx0U8ntdC1Cqy5ytPeqDzvCVQj1kk2dfAOe+k9aPCBTiwuX71tF28bi8vTQkSWAgxIY4IHVgy2Ej5EQcqkLPVht9G/JWMNhZ8t6PzWYnHqTdJDhoh5HklfnTMc5Amdf14ci+cM0Bf6R1uqmnurMxYixhGR+jAJyWPXsuIocMsviirl9HNsllALzGG4YVvygDmQh6ekha3ZBINxsShAVwQjbROAq1HsHMYTRBQclto+L+gkO5EFmhTuwIDAQAB"
//	if _, err := ncrypto.DecodeRSAPKCS1PublicKey(ncrypto.FormatPKCS1PublicKey(key)); err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestDecodeRSAPKCS1PrivateKey2(t *testing.T) {
//	var key = `-----BEGIN RSA PRIVATE KEY-----
//MIIEpAIBAAKCAQEA+AoH2rRQBgfu4yz7lUIhAjhMoCKi+sMks7QGd4tC/pdh8cUx
//9c5zhiRDsxP5uQ2b2h1iQ8SyVBX63N8sJcJlIPa2nz1q+hlpx0U8ntdC1Cqy5ytP
//eqDzvCVQj1kk2dfAOe+k9aPCBTiwuX71tF28bi8vTQkSWAgxIY4IHVgy2Ej5EQcq
//kLPVht9G/JWMNhZ8t6PzWYnHqTdJDhoh5HklfnTMc5Amdf14ci+cM0Bf6R1uqmnu
//rMxYixhGR+jAJyWPXsuIocMsviirl9HNsllALzGG4YVvygDmQh6ekha3ZBINxsSh
//AVwQjbROAq1HsHMYTRBQclto+L+gkO5EFmhTuwIDAQABAoIBAQDHaMkHuw7aymXz
//naZlSNTgxAJWe4Irt9c54fGRChTCOVI03SRdSaz2mxESV2zcqWQy5oypjukPUNxt
//1b3YXYCQD8iiHO50QxoXakJiMQ05JVrPHrK3E87f49/wXuqReDLCYTdthEnniLwL
//ZVaDuxhSPPz6IeDthSh6oYVEkjPOTMc7mLeU8GMbAWp26EaOePK/IFonttsg+CdS
//pzlgACLie1d2cz9BZVH8HxxjsvZI23wdaMwuO911bLXrNOTaJkv8VtnbXzWPAuPL
//CMT8HQD+NjqorY226aHm1v0dGM2qZh2kafiZjoKcwJswUq9u0B7YElrbHrxaZqhh
//+MBjmoYBAoGBAP8jTAZ10oZy6Pocxg4US7f0BXqtwuj/PYRuuiUp9WAs/u4FLJrR
//D0IJlOinEHhCDN6vprTxkbxv7ppVm+3G3+zZcAhNVPx99aH9okux4R+Kt1uHW04Y
//YiiB/DiOZZjJXiFrmSv+8/F1DCIGI8E8QLKhMRBsTlLw0ihlgHPeR8IxAoGBAPjg
//l9TSJ2RS1RQVSGRKEGsXBQtALcjbyp3u1T1LxAxcRmYUd/arq/YIRz6IxQpsSL6o
//kR0ql5vJ874Fhvi09IWCxzhA35uKokVRNCwTHgpGw78ke85VihjcO7kFh2X0Za1V
//xMaynDnwmtbvBUoery6rPovBXAjFSguGJZSwBy2rAoGAVfdMrHiSoN6jBXMggKny
//AV9lYTqfC+bEkZTFf7Qti2Rd2VFBAEEkxPl5hRKViA+pSncg0qL6meWVHcdurKVv
//2dxB1WpJEnJWi2hsgNjEo94Xfe2TpDKvEDsMAvpn3R0sbFJW3+4dv+PSXLx426kz
//m3yYjZIQqf2esiHRMDlbMGECgYEAqQzFlUX6g+bP3YhnV9gHNva38mBxuRGym3xz
//U+N3E35KQ2R76A3MDJ6q5gs47E+JShhuGYWo3aVb7yMJ/z3LKjUe5VMLkDem6bay
//3Y1OxmQy6k752bm4yBIwbwkWY2aS+h5cPytret+4DW7mBOa6z9R586wvu7VnClTg
//ihLXjWkCgYBnYK8a+iMWbYUBRZCxKY4TvENvCo/vakgkQUV5XCdZe7L4ce9/GT6N
//eWVqiNsxYH9XTgeP04U5aIRPr3rWOk+GcNgSx42XWW4wyJNKRO2BYooOnGPCM5Gi
//rtMCf50mpr6YvbEUkB7nqHaVkGMXy9MGZwX3Nzl8e53rff5FmzJXVQ==
//-----END RSA PRIVATE KEY-----
//`
//	if _, err := ncrypto.DecodeRSAPKCS1PrivateKey([]byte(key)); err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestDecodeRSAPKCS1PublicKey2(t *testing.T) {
//	var key = `-----BEGIN RSA PUBLIC KEY-----
//MIIBCgKCAQEA+AoH2rRQBgfu4yz7lUIhAjhMoCKi+sMks7QGd4tC/pdh8cUx9c5z
//hiRDsxP5uQ2b2h1iQ8SyVBX63N8sJcJlIPa2nz1q+hlpx0U8ntdC1Cqy5ytPeqDz
//vCVQj1kk2dfAOe+k9aPCBTiwuX71tF28bi8vTQkSWAgxIY4IHVgy2Ej5EQcqkLPV
//ht9G/JWMNhZ8t6PzWYnHqTdJDhoh5HklfnTMc5Amdf14ci+cM0Bf6R1uqmnurMxY
//ixhGR+jAJyWPXsuIocMsviirl9HNsllALzGG4YVvygDmQh6ekha3ZBINxsShAVwQ
//jbROAq1HsHMYTRBQclto+L+gkO5EFmhTuwIDAQAB
//-----END RSA PUBLIC KEY-----
//`
//	if _, err := ncrypto.DecodeRSAPKCS1PublicKey([]byte(key)); err != nil {
//		t.Error(err)
//	}
//}
//
//func TestDecodeRSAPKCS1PrivateKey3(t *testing.T) {
//	var key = `-----BEGIN PRIVATE KEY-----
//MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDdqW3b5Ms4yu0u
//0PMEw2Kj975WVu4nvRDrabtEA8+7tLg+LSZHtXdS4b/QM+Jfyk0wWuRlXGEaKELc
//fPGXcMpES7xQd5r83FspMgU7TZKVznEt+wRwAoUB8zgnpTTE0bBqRBV86JazIzSW
//SrE3DHvEl8QkufxEsUKAWE4lHMZyer44sHT1BSbPOPYKSSF1n4shGXsJnbWFFiTO
//MdYCu8sZK+sjn1B4QGZir9LZDxUwuW1cW+H3Q3DUfh2Gg6gt5AiQ1knjhwnvAoOB
//Ux/DBA7YTxOW2xP9bks3OJRRGhniAnSCgEeeHxyVvglFlHutpYmdWeQUxBchLC1N
//aDWvOZ67AgMBAAECggEAYTSf3MKk7C7xfIpdSEV2yHkomyxcHEkpVlNzsAwL2UET
//WxDprKDDxzEAsQlPyLoqx33Uky7D1ni5eX/BYqwvx2t/Fbqd2S25PIOXD76g/dke
//gQ+HjqdMNOK4SYqFbZqXRsK8uZ42GDDkKtRuirmKvVfgyEpXerhRdX9EJbpMvH6O
//V76GljSzaP7hlONQnqJwwiLbvYft8V5M7Dg396ZdHwCmpXcnUNZi2mOW11jau9jt
//32HtucQDDokDBaweAVTMjMjUc6JPpev5IxZwxxoSt5FwxUMFVSAojbW97IvuE4zX
//GsTmQKgJjz5b0/g2p/BJ9akOQJXd9VSDT2H1qnN+4QKBgQDz4IS+RwpcK8dUm2V3
//i2ffL8H9/6e3fEV9ddPwq16zntfKinTGMw3L/jHSEosWTbCtsey8YF20mt1QzE/Y
//UafaNq7SA7ynds0azidCX5Opnhto3jR7+bVwHEG+aM41A+O2HSuQkLRvnep3MQtu
//mPIlNBMUVhMQhpA1HtvnoAAY3wKBgQDorjV9s3kVsf78k4I9DG4wzRL0pQZQwVbs
//BFdx+SNB+NWOP3hQH2EaS+vo8td+Su8COCvVEaBkcj5E/VYpdM/+BimYVCsgO1IZ
//bpZ4Zl70zP7A7iIHJHz7nmHMqK5HtZVCqDU+rl8lwmsKhhZMv/uGkBYSG8qDmbJN
//utjK8C5JpQKBgQCY7j0vuQrFG+rmon/xhCkB60fiSNRamoJVF3NtY3mCd+RoLYTm
//2RMYzfFeA6TWuqdFdOK4ilGYtNh+F3EmfKYej9X2+sLt3PDnk0lJNPg7EFut2lHd
//QIzHneCYT86nriW4iPhNsoCD0eo7DjVWCHEoZUxln0hbP5bkyNIroMkOzwKBgQDO
//+rkGmnIpxG6oy6bdG4Oh0ar9CRd3GqVxyeLntdGRTALcSMWDsIY2WDcAv5TF+W5n
//MQR8wvNEpwORZzIQKqMiuuShZq3+PAaqLN+ZNqddaEVt5edLJ0AurBMfdSWjEHo4
//E11Z5s8ozatebyzDFTxv1Rhs/1/EdSJRVtamn8TtKQKBgQCTGpdbQdwE6DZY7oSV
//Kkhi1gbpbdnGmNgaBQDFBtEvzkQzXratuNIEogxuo67GoMVIdHEBFjJboPhNaVbm
//u30TPkduUxFgOQnA4qJNg65NpAYUGhMTcnqg4VIYpm4L4sMOSaToAcgdiwQIoqp4
//non9hCkG8Y1nmHTFIL52SS8LoA==
//-----END PRIVATE KEY-----
//`
//	if _, err := ncrypto.DecodeRSAPKCS1PrivateKey([]byte(key)); err == nil {
//		t.Fatal("解析 RSA Private Key 应该失败")
//	}
//}
//
//func TestDecodeRSAPKCS1PublicKey3(t *testing.T) {
//	var key = `-----BEGIN PUBLIC KEY-----
//MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3alt2+TLOMrtLtDzBMNi
//o/e+VlbuJ70Q62m7RAPPu7S4Pi0mR7V3UuG/0DPiX8pNMFrkZVxhGihC3Hzxl3DK
//REu8UHea/NxbKTIFO02Slc5xLfsEcAKFAfM4J6U0xNGwakQVfOiWsyM0lkqxNwx7
//xJfEJLn8RLFCgFhOJRzGcnq+OLB09QUmzzj2CkkhdZ+LIRl7CZ21hRYkzjHWArvL
//GSvrI59QeEBmYq/S2Q8VMLltXFvh90Nw1H4dhoOoLeQIkNZJ44cJ7wKDgVMfwwQO
//2E8TltsT/W5LNziUURoZ4gJ0goBHnh8clb4JRZR7raWJnVnkFMQXISwtTWg1rzme
//uwIDAQAB
//-----END PUBLIC KEY-----
//`
//	if _, err := ncrypto.DecodeRSAPKCS1PublicKey([]byte(key)); err == nil {
//		t.Fatal("解析 RSA Public Key 应该失败")
//	}
//}
