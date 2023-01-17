package validator

import (
	"errors"
	"go2ban/pkg/config"
	"go2ban/pkg/osUtil"
	"net"
	"regexp"
)

func CheckIp(target string) (end string, err error) {
	//Первый вариант поиска в строке ip адреса
	//target = regexp.MustCompile(`((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}`).FindString(target)
	target = FindIpOnByte(target)

	if target != "" && target[0] != '0' {
		if target[len(target)-1] == '0' && target[len(target)-2] == '.' {
			return "", err
		}

		whiteAddress := osUtil.GetLocalIPs()
		whiteAddress = append(whiteAddress, config.Get().WhiteList...)
		err = errors.New("This is white ip: " + target)
		for _, addr := range whiteAddress {
			if addr[len(addr)-1] == '*' {
				rexp := regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?){3}`)
				if rexp.FindString(addr) == rexp.FindString(target) {
					return "", err
				}
			}
			if target == addr {
				return "", err
			}
		}

		return target, nil

	}
	return "", errors.New("Wrong ip")
}

/*
Это функция FindIpOnByte, которая пытается найти IP-адрес в заданной строке и вернуть его в виде строки.
Функция использует массив байтов (buf) для хранения IP-адреса в процессе его анализа и
несколько переменных (C, I, P) для отслеживания текущей позиции в массиве байтов и во входной строке.
Функция перебирает входную строку, символ за символом, и использует серию условных операторов,
чтобы проверить, является ли текущий символ допустимой частью IP-адреса. Если допустимый символ найден,
он добавляется в массив байтов (buf), и переменные соответствующим образом обновляются.
Если найден недопустимый символ, функция проверяет, был ли найден полный IP-адрес (проверяя,
равно ли количество точек 3 и количество допустимых символов больше 5),
и если это так, она использует net.ParseIP для анализа IP-адреса и преобразования его в тип net.IP.
Затем используется метод .String() для возврата IP-адреса в виде строки.
Функция возвращает пустую строку, если во входной строке не найден допустимый IP-адрес.
*/
func FindIpOnByte(sts string) string {
	buf := make([]byte, 16)
	var C, I, P int8
	// Собираем буфер
	for _, b := range []byte(sts + " ") { // Пробел для последней итерации
		point := (b == '.')
		if ('0' <= b && b <= '9') && (C < 15) || point {
			if point {
				P++
				I = 0
			}
			if I <= 3 { // Блок из 3х чисел
				if !(C == 0 && point) { // Не первая точка
					buf[C] = b
				} else { // "Неверное вхождение" переписываем
					C = -1
					P = 0
				}
			} else { // "Неверное вхождение" переписываем
				C = -1
				P = 0
			}

			C++ // Итерация позиции в буфере
			I++ // Итерация позиции в блоке из 3х чисел
		} else {
			if (P == 3) && (C > 5) {
				//ParseIP приводит к четырем байтам и проверяет - !(n > 0xFF)
				ip := net.ParseIP(string(buf[:C]))
				if ip != nil {

					return ip.String()
				} else {
					return ""
				}
			}

			C, I, P = 0, 0, 0
		}
	}
	return ""
}
