package main

import "net/http"

var Chrome = http.Request{
	Method: "GET",
	Header: map[string][]string{
		`Connection`:                {`keep-alive`},
		`Cache-Control`:             {`max-age=0`},
		`sec-ch-ua`:                 {`" Not A;Brand";v="99", "Chromium";v="98", "Google Chrome";v="98"`},
		`sec-ch-ua-mobile`:          {`?0`},
		`sec-ch-ua-platform`:        {`"Linux"`},
		`Upgrade-Insecure-Requests`: {`1`},
		`User-Agent`:                {`Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36`},
		`Accept`:                    {`text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`},
		`Sec-Fetch-Site`:            {`none`},
		`Sec-Fetch-Mode`:            {`navigate`},
		`Sec-Fetch-User`:            {`?1`},
		`Sec-Fetch-Dest`:            {`document`},
		`Accept-Language`:           {`en-US,en;q=0.9,ru;q=0.8,uk;q=0.7`},
	},
}

