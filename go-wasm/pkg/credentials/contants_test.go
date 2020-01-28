package credentials

const (
	mnemonic = "opera initial unknown sign minimum sadness crane worth attract ginger category discover"
)

var (
	byteAttester             = []byte(`{"PrivateKey":{"XMLName":{"Space":"","Local":""},"Counter":0,"ExpiryDate":1611668018,"P":"v2Hw+dyLrmdzOpZQFL5OQVKFLXhizbpnhF6Xf0NGZ7h0OFtqN48NjdB8OAvxHAatq2vWd2YYXV7sPmeiVb9t1w==","Q":"zyspZ3jTEIslOAgH4VgMLh2YFU/6B2UHYrl3KYms8t+TLoylBicwiNDN0WBByqxOVtC5hXa/Oo51YSXnbnjZiw==","PPrime":"X7D4fO5F1zO5nUsoCl8nIKlClrwxZt0zwi9Lv6GjM9w6HC21G8eGxug+HAX4jgNW1bXrO7MMLq92HzPRKt+26w==","QPrime":"Z5WUs7xpiEWSnAQD8KwGFw7MCqf9A7KDsVy7lMTWeW/Jl0ZSgxOYRGhm6LAg5VYnK2hcwrtfnUc6sJLztzxsxQ==","ECDSA":"MHcCAQEEIJo4PrDs5A/jQigbeG19ygjU5u1jCcZuA/O84LEW00CZoAoGCCqGSM49AwEHoUQDQgAEiypa79vRgt6HSIMr8H3LMjjHGSxpwNwmVz3anmx25JMh5+CFyfz9HrOeyKxbTQpUpJ+jkIngQLyScv5aVvrpZw=="},"PublicKey":{"XMLName":{"Space":"","Local":""},"Counter":0,"ExpiryDate":1611668018,"N":"muB2QY2gqWLvETkMPhj/9XMIedZ75Umw4JcgoSRAqFVJpxYhc5b3P1NOOwQ0feT/4RLAxwoR3uT5AJN7VZbzG74Eopa5DeTl9UzHf+2Rmem1Hksdb6TjUV93kKVayqaTw1vhndzbYKBd3QsAqlKAdWna51hH2U5xbL6CwrHT4r0=","Z":"Q+SGUzUTWdVyqXDSHv04dRa9Vh1Tb0Ph9/K+CnyS5Kv5iiClZ8hFB3RXvtaikQR3fOI5nw94st8Zv0k1ZFfJpqidSp8SH38fk09AypZGsR8aeUw4BAye3dOIp8D066esLYUfVsdc0C1gEngbPdKh8zug74qgD28I6W79XKHd3Ds=","S":"dT/q9zhifzRtvri718xfJmjT48LsAAzTArEv/bZAAmppclgbvcwbraflndEeygU3Ld9YX0oFUqM3iXYXUbKBVj5UOg2lx2s2K0mWziZoaYs00lk9EfR9RYyNS6+HaGA3gUqEYjtvXZvEolD1gsDjy+k1eu5TieFwF6uehR2P3hU=","G":"a23TBsmnA1IwlGkM6TPmYXkpVVt3v2TC4rad2EbujlImqbvgrc4x4i28uapqOjmVERwvdR8+ScKPZzsJTrDy8uZzvRVumHHproZIQafvnRhp6gcH4oV+QSXbQYp4L6IL0sQwXlf6oVwPc7SlYQ3pm6tB21yNxyZt4tz/A/L5e+c=","H":"VrPYDlINT7c0fdc3OaWlrReL1uiM8RNqgR3sA2TtWIGT8pSWFrPhsba6I8FENLBLimcGw9WBfAnjjMOXI1Hfz/desXHU3qquNx9ivlXjHwXEvZUZGS/1fw5SDDCsGryTbO2q1SO4CTAIubfkra0GTOo5csEOBlMWCyFLORGlbdI=","T":"hNy7JhH2GS07vACZHUouyEQPSByjCHtBRWQkQgnMmLUAfx2n4FFHjmDekrdS6SUsmL0S6Rg/0PhGtiDijliLSIwuT3PAg095PWMo+FPh/Wohz3oClMuB9R4mR7la1TC1YBQUXynLSCfdCpn+izpyR1iYT9wZuTGxEE/ZkHwlyHE=","R":["HJgI2I+kyJEQ4KzxmgI/B0kLZzhAl/LXZsFy4ex/Axcy1nOD9dR8iOuOVVhrdYQLc34Zuth7AZ+Xd62lEwnXgW7So2FKWebqXLfaNfCJ+6HHSZZ8dzxp5pOEIVqLipL8TWUUHBly9EIvFeaIz2yXk7ebqQjN+Tvh/Ov3j/Qw+VU=","NqtFE/QgJdgtDLz+vAPX166F/mKz3g53Seylwfeh0hfi9lCGsVPWhcP1kBRI1FWsBIhEtDHUvB20wwYzNlPM6UM3Ie4OvXNUzaXH0y7qavkVBmYKbPMJdXaeqWF+1pNfDTSFAb59cmaWVtAPN3SD3sa3lto8U6oIEkuPAV1A2rI=","kL1mJ+vN0VBeihHPK75oNcae0rT7Js5qnC9FKpJAgqsYfxjKiGjWnqtUwfTPLDROYMHvEx20Hv45mFPORD2kHVdMSyisCYGecukgIUBxyP1IwhTk6+yEchOGUSaf4jgjZokflDDE77GxFveKFfvw0TAJUpqC0GvlaoR4rCN4pPE=","gwdEIS1l7dsW+mULFt2pOGmJ4zBJZKL981dgHrD/DTfCZ1j3lFbfm+M1bkfbBrSMQOFH2gMgXk0XDacpWdZS23i68viEgzyovhDVweUhDEw2WpDHrvUVbIYUNkcBM6ODfDMekBlaOnACVo5bUNAvyruYVuIAJSZTHNnN4iXSn88=","jaf/0rNUJiglvt9L9bSQdxK1nlHk8kdD5I5PpFGaipzuqjCjTQmSVp1P68DmW6DwbBBXVNQis4bFGIUMuoU4iBuwTdHQdghymxCUDfD0DVs/QQ1u3LsZeDmo8+h6kVutzr9IY4Ny+WATlPgkxiwrxJapbvFb7byQYN+o75MXdHc=","J5zio/wqvkZJ3kE6OHp8XTywI5HIDW6N4yWzDFUrRNeLHcsU+TSxfackXUIHsGJl1DD+n6CfYazuyrxIQVZvhbXY12ZlgX1+JUjpYC3qQpTNO588fiwot7anHXHMFxbbdggWdWKqF95yzhXs6pOU4jhcmugBtpu8uNxooyHA0/o="],"EpochLength":432000,"Params":{"LePrime":120,"Lh":256,"Lm":256,"Ln":1024,"Lstatzk":80,"Le":597,"LeCommit":456,"LmCommit":592,"LRA":1104,"LsCommit":593,"Lv":1700,"LvCommit":2036,"LvPrime":1104,"LvPrimeCommit":1440},"Issuer":"","ECDSA":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEiypa79vRgt6HSIMr8H3LMjjHGSxpwNwmVz3anmx25JMh5+CFyfz9HrOeyKxbTQpUpJ+jkIngQLyScv5aVvrpZw=="}}`)
	byteUpdate               = []byte(`{"sacc":{"data":"omNNc2dYu6NiTnVYgHBXNuh0hKvcL8wYEwqFsjaiRt3QhAorkWxmMnnaTnugAGwbmDuAOsre0gMKkXw+JXFQm3kWUnAIzcA71NU3yGs8giMPbGu/JLzZN8K/ZKv4Sf8mdFfL1Hh5N/GACk/eCW7uX7w1JwUXzR0+z8ZAp4mrpwSQJhvJRvmuQaRrhPslZUluZGV4AWlFdmVudEhhc2hYIhIg542kR7yDoMEFhYbUEni4Da0kZK3162KvG7Gb79HlbbtjU2lnWEcwRQIhANT3BdUZvFCbrMMRv1m0MTMDwBdeZ1bQeOyDOIIs+jTwAiA9elGZWkmNbUBiCyUGNdBCyeOSltiVLCcWO3lL/I6+Mw==","pk":0},"i":1,"hash":"EiAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA==","e":[""]}`)
	byteClaimer              = []byte(`{"MasterSecret":"onIfyirb1JJTKOvFc4qMov6os2UdxyQwdRMkQl0Va5c="}`)
	byteAttesterSession      = []byte(`{"context":"pBwrpzMiQlnWgH4b5SJ2C+w66AGmg+ii6dG37sKRvVA=","nonce":"GLxG+BbGZXMYXg=="}`)
	byteInitiatAttestation   = []byte(`{"nonce":"GLxG+BbGZXMYXg==","context":"pBwrpzMiQlnWgH4b5SJ2C+w66AGmg+ii6dG37sKRvVA="}`)
	byteAttestationRequest   = []byte(`{"commitMsg":{"U":"Vemzabzs8/G9g/cw/3WXwZP7tGcTJ93KSo/9gB7bpzeL8Wol0pTCFtXK93snFfJRjVfHAWA0O8Spi259jB2XUi+OBTNGL6gZ04uBcMR3mrJR4/aSw8rFM/Ds5WY1oE6g5Sc2C8mSJ6Azz6Pmku5XDZojtH0QYDCmgQSd45ehnTg=","n_2":"Zx2Q1fmkA1so9w==","combinedProofs":[{"U":"Vemzabzs8/G9g/cw/3WXwZP7tGcTJ93KSo/9gB7bpzeL8Wol0pTCFtXK93snFfJRjVfHAWA0O8Spi259jB2XUi+OBTNGL6gZ04uBcMR3mrJR4/aSw8rFM/Ds5WY1oE6g5Sc2C8mSJ6Azz6Pmku5XDZojtH0QYDCmgQSd45ehnTg=","c":"ZnA9FQdKIjZ4oPIkY7GhssJcKmyb/NYJ3gtUnts38Qo=","v_prime_response":"YjdRvLcho/uH+cbhK0P4YPw/uL4Dmp2Z2SanfBdoO2ZnQnzZpZpLbaFwBvNtRTzHz6XadiCWOo21UwV0lvZcu1uaePOXmhys3W99tDsJGkmw8biPqgMmJc4CcIiQ1Bd7RShgVoa6aCJvcYt9N/DKqDhPkSv536BZKti/FFHqTUqRT1kxRkyqmyK5c/20HwkLyVYzUYG8ukMjkk9N42eZgr5g8lVxcXqyM2KgrrwwkG4vZom0","s_response":"7e6CIBoOlreH2RE2wrCFtxIbCeau6ZgH6FjBN2tElEBYHCgIpCtBrHaJdVln+w5cxOAlkC42AwGYHb2V3D38Q6GTRfL1xgLR87Q="}]},"claim":{"contents":{"age":34,"gender":"female","name":[{"a":1,"b":2},2,3],"special":true},"ctype":"0xDEADBEEFCOFEE"}}`)
	byteAttestClaimerSession = []byte(`{"cb":{"Secret":"onIfyirb1JJTKOvFc4qMov6os2UdxyQwdRMkQl0Va5c=","VPrime":"qCjuGtgeMGZxyHOckyDBOb7gNy8C6EjYsZwauvxvrj3MURHtUXvgl/dPWmwUvD/J16WBSN2Ej0CIZyPR2Rxk075LEDC5HGyg7TghdCvq568Uxtdb/1Bvl+mqowLCT9hqfY/NX6Uf/NqdCw1QJwTYNRsCHWhRP1+Oss/H4uktiW2Z6CtdhNU2PRLV","VPrimeCommit":null,"Nonce2":"Zx2Q1fmkA1so9w==","U":"Vemzabzs8/G9g/cw/3WXwZP7tGcTJ93KSo/9gB7bpzeL8Wol0pTCFtXK93snFfJRjVfHAWA0O8Spi259jB2XUi+OBTNGL6gZ04uBcMR3mrJR4/aSw8rFM/Ds5WY1oE6g5Sc2C8mSJ6Azz6Pmku5XDZojtH0QYDCmgQSd45ehnTg=","UCommit":"AQ==","SkRandomizer":null,"Pk":{"XMLName":{"Space":"","Local":""},"Counter":0,"ExpiryDate":1611668018,"N":"muB2QY2gqWLvETkMPhj/9XMIedZ75Umw4JcgoSRAqFVJpxYhc5b3P1NOOwQ0feT/4RLAxwoR3uT5AJN7VZbzG74Eopa5DeTl9UzHf+2Rmem1Hksdb6TjUV93kKVayqaTw1vhndzbYKBd3QsAqlKAdWna51hH2U5xbL6CwrHT4r0=","Z":"Q+SGUzUTWdVyqXDSHv04dRa9Vh1Tb0Ph9/K+CnyS5Kv5iiClZ8hFB3RXvtaikQR3fOI5nw94st8Zv0k1ZFfJpqidSp8SH38fk09AypZGsR8aeUw4BAye3dOIp8D066esLYUfVsdc0C1gEngbPdKh8zug74qgD28I6W79XKHd3Ds=","S":"dT/q9zhifzRtvri718xfJmjT48LsAAzTArEv/bZAAmppclgbvcwbraflndEeygU3Ld9YX0oFUqM3iXYXUbKBVj5UOg2lx2s2K0mWziZoaYs00lk9EfR9RYyNS6+HaGA3gUqEYjtvXZvEolD1gsDjy+k1eu5TieFwF6uehR2P3hU=","G":"a23TBsmnA1IwlGkM6TPmYXkpVVt3v2TC4rad2EbujlImqbvgrc4x4i28uapqOjmVERwvdR8+ScKPZzsJTrDy8uZzvRVumHHproZIQafvnRhp6gcH4oV+QSXbQYp4L6IL0sQwXlf6oVwPc7SlYQ3pm6tB21yNxyZt4tz/A/L5e+c=","H":"VrPYDlINT7c0fdc3OaWlrReL1uiM8RNqgR3sA2TtWIGT8pSWFrPhsba6I8FENLBLimcGw9WBfAnjjMOXI1Hfz/desXHU3qquNx9ivlXjHwXEvZUZGS/1fw5SDDCsGryTbO2q1SO4CTAIubfkra0GTOo5csEOBlMWCyFLORGlbdI=","T":"hNy7JhH2GS07vACZHUouyEQPSByjCHtBRWQkQgnMmLUAfx2n4FFHjmDekrdS6SUsmL0S6Rg/0PhGtiDijliLSIwuT3PAg095PWMo+FPh/Wohz3oClMuB9R4mR7la1TC1YBQUXynLSCfdCpn+izpyR1iYT9wZuTGxEE/ZkHwlyHE=","R":["HJgI2I+kyJEQ4KzxmgI/B0kLZzhAl/LXZsFy4ex/Axcy1nOD9dR8iOuOVVhrdYQLc34Zuth7AZ+Xd62lEwnXgW7So2FKWebqXLfaNfCJ+6HHSZZ8dzxp5pOEIVqLipL8TWUUHBly9EIvFeaIz2yXk7ebqQjN+Tvh/Ov3j/Qw+VU=","NqtFE/QgJdgtDLz+vAPX166F/mKz3g53Seylwfeh0hfi9lCGsVPWhcP1kBRI1FWsBIhEtDHUvB20wwYzNlPM6UM3Ie4OvXNUzaXH0y7qavkVBmYKbPMJdXaeqWF+1pNfDTSFAb59cmaWVtAPN3SD3sa3lto8U6oIEkuPAV1A2rI=","kL1mJ+vN0VBeihHPK75oNcae0rT7Js5qnC9FKpJAgqsYfxjKiGjWnqtUwfTPLDROYMHvEx20Hv45mFPORD2kHVdMSyisCYGecukgIUBxyP1IwhTk6+yEchOGUSaf4jgjZokflDDE77GxFveKFfvw0TAJUpqC0GvlaoR4rCN4pPE=","gwdEIS1l7dsW+mULFt2pOGmJ4zBJZKL981dgHrD/DTfCZ1j3lFbfm+M1bkfbBrSMQOFH2gMgXk0XDacpWdZS23i68viEgzyovhDVweUhDEw2WpDHrvUVbIYUNkcBM6ODfDMekBlaOnACVo5bUNAvyruYVuIAJSZTHNnN4iXSn88=","jaf/0rNUJiglvt9L9bSQdxK1nlHk8kdD5I5PpFGaipzuqjCjTQmSVp1P68DmW6DwbBBXVNQis4bFGIUMuoU4iBuwTdHQdghymxCUDfD0DVs/QQ1u3LsZeDmo8+h6kVutzr9IY4Ny+WATlPgkxiwrxJapbvFb7byQYN+o75MXdHc=","J5zio/wqvkZJ3kE6OHp8XTywI5HIDW6N4yWzDFUrRNeLHcsU+TSxfackXUIHsGJl1DD+n6CfYazuyrxIQVZvhbXY12ZlgX1+JUjpYC3qQpTNO588fiwot7anHXHMFxbbdggWdWKqF95yzhXs6pOU4jhcmugBtpu8uNxooyHA0/o="],"EpochLength":432000,"Params":{"LePrime":120,"Lh":256,"Lm":256,"Ln":1024,"Lstatzk":80,"Le":597,"LeCommit":456,"LmCommit":592,"LRA":1104,"LsCommit":593,"Lv":1700,"LvCommit":2036,"LvPrime":1104,"LvPrimeCommit":1440},"Issuer":"","ECDSA":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEiypa79vRgt6HSIMr8H3LMjjHGSxpwNwmVz3anmx25JMh5+CFyfz9HrOeyKxbTQpUpJ+jkIngQLyScv5aVvrpZw=="},"Context":"pBwrpzMiQlnWgH4b5SJ2C+w66AGmg+ii6dG37sKRvVA=","ProofPcomm":null},"claim":{"contents":{"age":34,"gender":"female","name":[{"a":1,"b":2},2,3],"special":true},"ctype":"0xDEADBEEFCOFEE"}}`)
	byteAttestationResponse  = []byte(`{"proof":{"c":"GKTset5GvAwipmYQNyYGUCQIK0qO3SdEC3k6QcZP9J8=","e_response":"HxY672/wBmOgfz4UMcHT89iNWGtK6uQLmus1pX2w//bNKa0teqOHsqsprZZ3gotqMI+0q79YHP2ZEhhpEZd0JUS2KwkRNkX8AszeT79e9BywAcMA4qKsbRzRh9t0NxbxePnCp/mmJFOYjlSnqzcYl+BgRx7Ghfj9CvGQfS2rY3M="},"signature":{"A":"aCKRn2tSJ0P/fXln4ptxreSFe0BMkE1zP0SrGjY0wOMmMSOfhOW2WyFDAKLUPbwMrs2TU81Pgfwko8tVUdb07H0dAewsiFKgvpEvi64EO7ot1q2DfnClzvOh7AMbCFyrG07+ZmywR3KdqkaB+/uPRtdimrnF6+dmiHqVHivZ/6c=","e":"EAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAN/BVLoyT9wBY8FcUHO5h","v":"D4d8lREAa4cIHj/qMZ8pX8WuTIOjansKetguswW4mP1QnN3ek65yAYOVqUOJtoIakhb72KPiEYB3MkzAwXdljsMcIHcfQ/pjPPvNPfzYcydg27GgmJkqqb41wsyQ/U7Jr2VtxZxFj8GiejXso14MvAnQR1DXDwRG3K41iebGeMFIqdyDNSmV4mPkDIIENdzhzroZ5bbJqCSMgZnDBfO3s6sDz/Vh5oFcrreejO9FFHajUhEE3II40cZxwpn3FgWsEI8Dj8qF6GCPiOvxl84Nb52e/uJE","KeyshareP":null},"nonrev":{"u":"GBtz58fg69gKTnU8+H3/NcS2e8dXgPU/PCEASy670DnRsqk45SiK3unM0UJauXAcBscfiss0UQR8d0oIJCH63/gD/75iyQSv1eAt/DcEd1GNL8qEqefCCQowOOJt8WoOsVwBJplTu5VRmQk/oxuPKuLSJ3KSsLJcDghva3lY5rc=","e":"BFIBJRK8fUS2sTJ+AISy4dTN3DMQ09kOpw==","sacc":{"data":"omNNc2dYu6NiTnVYgHBXNuh0hKvcL8wYEwqFsjaiRt3QhAorkWxmMnnaTnugAGwbmDuAOsre0gMKkXw+JXFQm3kWUnAIzcA71NU3yGs8giMPbGu/JLzZN8K/ZKv4Sf8mdFfL1Hh5N/GACk/eCW7uX7w1JwUXzR0+z8ZAp4mrpwSQJhvJRvmuQaRrhPslZUluZGV4AWlFdmVudEhhc2hYIhIg542kR7yDoMEFhYbUEni4Da0kZK3162KvG7Gb79HlbbtjU2lnWEcwRQIhANT3BdUZvFCbrMMRv1m0MTMDwBdeZ1bQeOyDOIIs+jTwAiA9elGZWkmNbUBiCyUGNdBCyeOSltiVLCcWO3lL/I6+Mw==","pk":0},"Updated":"0001-01-01T00:00:00Z"}}`)
	byteCredential           = []byte(`{"credential":{"signature":{"A":"aCKRn2tSJ0P/fXln4ptxreSFe0BMkE1zP0SrGjY0wOMmMSOfhOW2WyFDAKLUPbwMrs2TU81Pgfwko8tVUdb07H0dAewsiFKgvpEvi64EO7ot1q2DfnClzvOh7AMbCFyrG07+ZmywR3KdqkaB+/uPRtdimrnF6+dmiHqVHivZ/6c=","e":"EAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAN/BVLoyT9wBY8FcUHO5h","v":"D4d8lREAa4cIHj/qMZ8pX8WuTIOjansKetguswW4mP1QnN3ek65yAYOVqUOJtoIakhb72KPiEYB3MkzAwXdljsMcIHcfQ/pjPPvN5iXGjf9/DBgSYQzHPN72/ItxNH3Ml65GdzhgSr4SKHO49G/6DYWw30gmaXBbmO3/YYxHwZ7NOR0LnE1nu4BI4EBPRg2a6ya60u7rHFB3aUjXzMsTsvtzZ98MiYQe/pAJCn8SdBvDTuuh54+I+MtJ97T5M239T+6SQppNy0m9ElmLf/lq9HLVO/UZ","KeyshareP":null},"attributes":["onIfyirb1JJTKOvFc4qMov6os2UdxyQwdRMkQl0Va5c=","/wAAAAAAAAAMY29udGVudHMuYWdlAAAAAAAAAAVmbG9hdAAAAAAAAAAIQEEAAAAAAAA=","/wAAAAAAAAAPY29udGVudHMuZ2VuZGVyAAAAAAAAAAZzdHJpbmcAAAAAAAAABmZlbWFsZQ==","/wAAAAAAAAANY29udGVudHMubmFtZQAAAAAAAAAFYXJyYXkAAAAAAAAAFP9beyJhIjoxLCJiIjoyfSwyLDNd","/wAAAAAAAAAQY29udGVudHMuc3BlY2lhbAAAAAAAAAAEYm9vbAAAAAAAAAABAQ==","/wAAAAAAAAAFY3R5cGUAAAAAAAAABnN0cmluZwAAAAAAAAAPMHhERUFEQkVFRkNPRkVF"],"nonrevWitness":{"u":"GBtz58fg69gKTnU8+H3/NcS2e8dXgPU/PCEASy670DnRsqk45SiK3unM0UJauXAcBscfiss0UQR8d0oIJCH63/gD/75iyQSv1eAt/DcEd1GNL8qEqefCCQowOOJt8WoOsVwBJplTu5VRmQk/oxuPKuLSJ3KSsLJcDghva3lY5rc=","e":"BFIBJRK8fUS2sTJ+AISy4dTN3DMQ09kOpw==","sacc":{"data":"omNNc2dYu6NiTnVYgHBXNuh0hKvcL8wYEwqFsjaiRt3QhAorkWxmMnnaTnugAGwbmDuAOsre0gMKkXw+JXFQm3kWUnAIzcA71NU3yGs8giMPbGu/JLzZN8K/ZKv4Sf8mdFfL1Hh5N/GACk/eCW7uX7w1JwUXzR0+z8ZAp4mrpwSQJhvJRvmuQaRrhPslZUluZGV4AWlFdmVudEhhc2hYIhIg542kR7yDoMEFhYbUEni4Da0kZK3162KvG7Gb79HlbbtjU2lnWEcwRQIhANT3BdUZvFCbrMMRv1m0MTMDwBdeZ1bQeOyDOIIs+jTwAiA9elGZWkmNbUBiCyUGNdBCyeOSltiVLCcWO3lL/I6+Mw==","pk":0},"Updated":"0001-01-01T00:00:00Z"}},"claim":{"contents":{"age":34,"gender":"female","name":[{"a":1,"b":2},2,3],"special":true},"ctype":"0xDEADBEEFCOFEE"}}`)
	byteClaim                = []byte(`{"contents":{"age":34,"gender":"female","name":[{"a":1,"b":2},2,3],"special":true},"ctype":"0xDEADBEEFCOFEE"}`)
	byteVerifierSession      = []byte(`{"context":"QTfMVxfhxphPmO4+wu0uqJ3jbbm7alILhF3zu60KzQk=","nonce":"8YI73+2CsQOBvLg2oEEc1ZZfkaPYXytZIpyMfcwYbDI=","reqNonRevocationProof":true,"reqMinIndex":1}`)
	bytePresentationRequest  = []byte(`{"partialPresentationRequest":{"requestedAttributes":["contents.name","contents.age","contents.special","contents.gender"],"reqNonRevocationProof":true,"reqMinIndex":1},"context":"QTfMVxfhxphPmO4+wu0uqJ3jbbm7alILhF3zu60KzQk=","nonce":"8YI73+2CsQOBvLg2oEEc1ZZfkaPYXytZIpyMfcwYbDI="}`)
	bytePresentationResponse = []byte(`{"proof":{"c":"So0dKi/8e5C/6GS2kbmSDKzuwgRaHtH5An++2Bb+FWU=","A":"DLGrWtscMRUWMOTPGHTOEAe+GlzqqnRKbmmeyoxksXk0JS2BtdlQtgmRjpeGBpP4UX0iKNxFKYH7vD18kxKHO8Ras6lFnJb20l26syfbLohU4Ikma1IIeIv8U8Qpbert5OvubEcWSvQnhHpKaKWkfhtW3IsIxXhTVm47Uwmvctg=","e_response":"2eQNapLHrr/97E9mRYBcvQfq4jeiW5YkrukaQwnOs4mPDpaPpLYfSrHR4uQtC0MOio6lqY8bbPKm","v_response":"DQoOmH3bd27uVoL1qOM5nKEyzoBnLYpsGxF2G0U0tYVV/uUVVWt82gFfVrNBWW3aaNKwYZ7YepL2/2vLjr17oGw8iY1IEPJe6l6l59AYBz3TKg1Miuu+hY6Fe41ChJDA2tz1Ph+BJY7HQFKrs41aSbpOgnrAmaSKGM51OQoYVbZCQHW//Av8QHksMDtVe7+T80GiFeuSB72XQLC2DOT2i9FcCg5ZDoSeaa4rT7ofycb1LEqyQsE/XK08Io1yrMFvIAQ4NtF1LzP/Lyqrm8dF88JlsXOnXTNzG1u3RLeq88ppnZvZb8B9AESOIe6f9doLANpMG6Hemx56ICNNvZdq","a_responses":{"0":"dIapGI+IT9gk9mBUgxe6TgMz2I0ABfxoBOz/qCEcKNe4q7z4qr1iAW1k+Une6KY5Dc2kH9BCR2V/5tWEhkMqFS8F9mCTKz9w1ww=","5":"cStX8KsJoyuIQ0Pn+4pOvOxl+Twd7Fg1nl53BM9QggORGpvVDbzb9+119jkdopqxcYN9Z88Ikr7zakT7WiQI6OlcdDYHOv5FbHE="},"nonrev_response":"BvZ9y8pLeo6LgKpvtYCZ1zaouMqUtn4f5tEF8IGL7LdIBmgRmmdjAujOgiFA/cPRsDVur2lrXTYWVPmo3/5d5+4lNLdgenUOPQ==","nonrev_proof":{"C_r":"Nnw84OfV94ZhOQ1iMHZMyHhzmWQtuexxgIJY0WSua21/b78Hmqd7S8vlH4tVZaiJ2Nhenog7u4eWLFl/JOeKJEMT98UjvZnmXsENn1On+7TB+MeqZccjo27tE0kUp+2HUbX3T7Tv4Q5kZz2amtw7+nD3sjHZIVt3S1YBj5XZ03A=","C_u":"H/8aOD7KF6MCI00s+/dHS5zWhOSMuncxeG8zgLnKXRat/FLfKzLVzx5daMFgGWhi/z7lrq3sr5RmRcKb5+f+Vw7X1W56qwYcKYViR1XCPiZIyRPftnw5X/cbgL40RqQl9q61fxWGXwFs+HqZ/mPAzYVZf4fmKTalzei44bmbVHE=","responses":{"beta":"mkaLCG7hRk/K49H7OuO839gao6Dy770od0Pe/KsHQkELQUcvbNFzevevwmnmjtD/kaLt9vUKeulFNfP32coX1VPY8o5esC1XDsJh0GsEaIbG7f4sKT6U0mgcT8N0/kXoByiKxMk9S4jjkXJdtvcVy6WVcFNEupWM9xgX4m58Au5oR9Kx2mdZNeuYNbGTrYEASJzyYSH3j71vfEILt4WvEeO1U0Hlow46wvxtKA/nPbfcj9H1sN7hjB8WQldMuwcqujxCBWMIgMc=","delta":"dtIElA05GiZfmMLq3KmUeqat3X8QzLN2rcNyjiUd2PMoJpfrMxPOz+pUo6E5mzVhQzH7GEuNMiayeMtH/YN4NFej2XSPCw79nnW6Y9zlQSDkuPQvTV3AgWdDNAPSiBSHHEwE5RmtaCpjVcQ1eRHnSrN3kV+OPUZ/4zgXSPYGT2smlzVpncuVXzqwXqDpYs6A+rH+blvZ9q/edYzAyjfOsjT587Ioi7G6NMQMnB3/AbYmL8jFf3RUzqN5C96HtVwCW4cqQOn23I0=","epsilon":"F9OiNExZX2Wf1pv8XVpEPj+a+zN1szRDKJH5fH608/u5AxtyDRnkZiJUXJzt4JBCV4xVjU+QgVKKdKvImzRYJXsBVvAwmj1BwaG/npqu46AUjQYbmgD36gmbEpAkcydV67Ok4b8P339FetH7yISIq6TJE4ryZCgQsAsjli70QWkM+HaDlsMyoyJnYNh5UHHn4YvQ3+7ot4kCrSN8jIEpn1kvGQ7VZ516eEOmwpkgeLg=","zeta":"BC1gv0JT2sMzQPBYi2uzO+SMlsJsg2bDl758V13Tb44Xn2sJRn0MyKYp8EXa2x98zyI1L0H6qfqVy8c2Y53cDDm4at5HC6wYoGMoIXwojaja3cDZcz5QESdqXpJtCEQ0i4nbt/BNdg+9xXqJbcVOdeX1khEpNPkxFgfp7w/KALDkWUtbyXXi/OkF1uhi5+GuxXBdCNO+2BrNhblkUI2745gMMftkLQzAqYHm7WolkP8="},"sacc":{"data":"omNNc2dYu6NiTnVYgHBXNuh0hKvcL8wYEwqFsjaiRt3QhAorkWxmMnnaTnugAGwbmDuAOsre0gMKkXw+JXFQm3kWUnAIzcA71NU3yGs8giMPbGu/JLzZN8K/ZKv4Sf8mdFfL1Hh5N/GACk/eCW7uX7w1JwUXzR0+z8ZAp4mrpwSQJhvJRvmuQaRrhPslZUluZGV4AWlFdmVudEhhc2hYIhIg542kR7yDoMEFhYbUEni4Da0kZK3162KvG7Gb79HlbbtjU2lnWEcwRQIhANT3BdUZvFCbrMMRv1m0MTMDwBdeZ1bQeOyDOIIs+jTwAiA9elGZWkmNbUBiCyUGNdBCyeOSltiVLCcWO3lL/I6+Mw==","pk":0}},"a_disclosed":{"1":"/wAAAAAAAAAMY29udGVudHMuYWdlAAAAAAAAAAVmbG9hdAAAAAAAAAAIQEEAAAAAAAA=","2":"/wAAAAAAAAAPY29udGVudHMuZ2VuZGVyAAAAAAAAAAZzdHJpbmcAAAAAAAAABmZlbWFsZQ==","3":"/wAAAAAAAAANY29udGVudHMubmFtZQAAAAAAAAAFYXJyYXkAAAAAAAAAFP9beyJhIjoxLCJiIjoyfSwyLDNd","4":"/wAAAAAAAAAQY29udGVudHMuc3BlY2lhbAAAAAAAAAAEYm9vbAAAAAAAAAABAQ=="}}}`)
	bytePresentation         = []byte(`{"contents":{"age":34,"gender":"female","name":[{"a":1,"b":2},2,3],"special":true}}`)

	// Combined...
	byteCredential2              = []byte(`{"credential":{"signature":{"A":"DPrOkTzubCAdRTqccMO1nWrwFhmlJxTYmvKjnYh4wAFfq0mBFj8BqkOgb4p0MjUwxRNs3thDHeEICyMgPFLM6jlK4P9CrrdZSaFDpssRg6PKuqyVMphrSrPmVfNfDV8Vy3kOiD1mwchMHDEQ3ANpLD0az0mPSnmUOndv9Zcqw/E=","e":"EAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAY+wfXpfhFbE0IF36qGUf","v":"DzXzRICOxf2GoQ9Ite7kvw0JYFxe7nds7oYVfJ7csa3wRybfGQCrIcqf8ddZJNpxLyqLVfAHQdo2DXqLWxqA+wgwr8aefz5BBAd+aTqToLrkH9UW+OiWgJH4G2eBxOp9HyiWQfYQW4S3I+z8swxBIx2buRiI9kJxsKf6LBVUiZ9m/o8YldrMxolW/BI3Xo4IDN6mOeFRXNEOYQWj8HEyhhR2GSpjFYVH+ww6ItrWoJqxT2XYbpGp7KwZShHuOpirQ8/JYcSHlXLTTACZoeJQbju9+uHR","KeyshareP":null},"attributes":["onIfyirb1JJTKOvFc4qMov6os2UdxyQwdRMkQl0Va5c=","/wAAAAAAAAAMY29udGVudHMuYWdlAAAAAAAAAAVmbG9hdAAAAAAAAAAIQEEAAAAAAAA=","/wAAAAAAAAAVY29udGVudHMubGlrZWROdW1iZXJzAAAAAAAAAAVhcnJheQAAAAAAAAAI/1sxLDIsM10=","/wAAAAAAAAANY29udGVudHMubmFtZQAAAAAAAAAGc3RyaW5nAAAAAAAAAAVCZXJ0YQ==","/wAAAAAAAAAQY29udGVudHMuc3BlY2lhbAAAAAAAAAAEYm9vbAAAAAAAAAABAQ==","/wAAAAAAAAAFY3R5cGUAAAAAAAAABnN0cmluZwAAAAAAAAAPMHhERUFEQkVFRkNPRkVF"],"nonrevWitness":{"u":"VcPJAvIPjTvJWQh4ZZsjqyKH3k+kPIoYP6e2fmWF89sKodnD2slN99T0DWlg/nKDw+z6xAPYeCJ0k1wsOHfZUdF00P0CkHRxs/vapUf053nFfmAQ80U5yds0fcsApg0joyBjuYUdDbDC2m3lT8zqo0VuxH2ZTf+WwR9BkKaFAKc=","e":"BO2VVOUEeWypMB0s67YSKi4K1nTEIkCmpQ==","sacc":{"data":"omNNc2dYu6NiTnVYgHBXNuh0hKvcL8wYEwqFsjaiRt3QhAorkWxmMnnaTnugAGwbmDuAOsre0gMKkXw+JXFQm3kWUnAIzcA71NU3yGs8giMPbGu/JLzZN8K/ZKv4Sf8mdFfL1Hh5N/GACk/eCW7uX7w1JwUXzR0+z8ZAp4mrpwSQJhvJRvmuQaRrhPslZUluZGV4AWlFdmVudEhhc2hYIhIg542kR7yDoMEFhYbUEni4Da0kZK3162KvG7Gb79HlbbtjU2lnWEcwRQIhANT3BdUZvFCbrMMRv1m0MTMDwBdeZ1bQeOyDOIIs+jTwAiA9elGZWkmNbUBiCyUGNdBCyeOSltiVLCcWO3lL/I6+Mw==","pk":0},"Updated":"0001-01-01T00:00:00Z"}},"claim":{"contents":{"age":34,"likedNumbers":[1,2,3],"name":"Berta","special":true},"ctype":"0xDEADBEEFCOFEE"}}`)
	byteCombVerifierSession      = []byte(`{"context":"jdIxpFxHzsMz9Dxx+I8bsxM3eoafyoi/lRh7QDoE4xk=","nonce":"ffTPJqkvTNIofGaH/jNs7mv35FjpdjrnuitetxBfghk=","partialRequests":[{"requestedAttributes":["contents.age","contents.gender","contents.name","contents.special"],"reqNonRevocationProof":true,"reqMinIndex":1},{"requestedAttributes":["contents.name","contents.likedNumbers"],"reqNonRevocationProof":true,"reqMinIndex":1}]}`)
	byteCombPresentationRequest  = []byte(`{"partialPresentationRequests":[{"requestedAttributes":["contents.age","contents.gender","contents.name","contents.special"],"reqNonRevocationProof":true,"reqMinIndex":1},{"requestedAttributes":["contents.name","contents.likedNumbers"],"reqNonRevocationProof":true,"reqMinIndex":1}],"context":"jdIxpFxHzsMz9Dxx+I8bsxM3eoafyoi/lRh7QDoE4xk=","nonce":"ffTPJqkvTNIofGaH/jNs7mv35FjpdjrnuitetxBfghk="}`)
	byteCombPresentationResponse = []byte(`{"prooflist":[{"c":"vjAUx4iyASz33XNJiyRwle17UbLoMMH7DIxIWk4FGV4=","A":"R2yyEnobAhLen1cZ+zORiYjhqKQdz/Zzy6DI7q9S/iqwQk6SMwQ4tJPA7UGMA0Z+o6kYJK8txInW90Cp1BBELnQ8XoApkyU/Po2jUl+tVoP8aBivIMRnFTSJxxoACyvldEr7bR0M2dm1JU/L7xGlq6W3hKYFb5xQEkvG/cxHyug=","e_response":"nUN2KfIA2N7daJL59DYX7aOufMUYS/jUltknhF0+8Q3IJdFusevd7ChkAM9nwyz3xNElQTooImQz","v_response":"DRbdA/MoTpg1e49VAb/H1zi8RwW4hf3V38w/V+OHLH8UaOOorZkQM/cnxPFA0h+jxwBAHBeMUVtbLWjCsxKZoGvtsqNdTU39HhVdzs/Jp/atu9TCfwJMDQ3ZeQgBD7qJp0JkbNMMuAjDRMLi+xMXvRKyyfzuvS8tllu4D+JKNAhpZe2ADYs4en6s1O2IgAhcy72oJXpVOSDerqRmkXeiRGSVYROkfNUQuXsJancFpzththcBjizUecXmmvj4TK3lSOp6XJYJH/+fV8c+AtUju+5ckeILF8tomqhImuk5w8a2ahoTfiECoP70iJ2h/LD49aLjHYzx+VqbG8AjBHBR","a_responses":{"0":"P8LzPt8+aQgwuDrOAM51h85UL3vUHS4YHTGAkzP+WvKz8e9VuAuyKvUUvv8NEadG5qFxTjNzDcvBnWVhwqxLjRMbQ0VSncfbJJw=","5":"E10EM++AQwwKWzanQH9Frlk9VJEsBg8ND7PRGXtKq8FvEaDwbfOc0fBt8Ct5jAzsZmNdzWzbIM3KFgS0H9Q7fDV82Ac7qO/tJjE="},"nonrev_response":"AWPwZO+umYUkWW1o5YJjIWagsRDV+e4135kym1iiuK9vn3ki4oW7CB9eJ1PMR6PubD35QTenGz9ybX7epa0+0O4dfLrL153Fnw==","nonrev_proof":{"C_r":"ISdUCAejBePwD+0SkgxTFu35nuqqFfM5/N8FKU4iSwH780BuBBGw5kiT5ISZvI3fuGfAaZbFLqnLAiMN6/fI9HLS7ji0QhV1oHoM5fyhJrb4Vg/ueQIM6sA73pgMyitlWQ6QGIcQAhZnTzOcpS4aanrR4KBxLg252ZkpG05x+Fc=","C_u":"YFVzRqYCacdQCWh3wpBmk0m+Py7KRbBz9LIRNwAiSzyTRjFjPMfvI6PvNtDHhuQrm4o1JbJfj/LQsbUStbIIyEjbdedC1osfpHRljGI6tZNzIog84ejBkTTu7+Th37mMddjXMTwRwxzdgNnC/5W1sSsjfSaC+x7ZNhax+yWWlk0=","responses":{"beta":"Du/rIp5qDmj6HHIVm46Vcsl3fIGnMo/JfJrFHGZEOcrtsAsxC/zOipr6s6vLtF1aweViWY/U/FCudfSJAr7to+cPzHpjZLCEsDDv/dchRLlmQSDD1m9i9iZEijIVw3ngEWM0tSUPvffsL3iyEo/MKEQS5hAZR689veW9UdbN43Bro5p8cCeqKKK/93Fm+S3qokHryPyp5NnkORQXUz+5TeqK+v9rWCKwH+hQTVPtnO1EHlji/EvED8vkheg/1HLUE3pJ0BZUA7I=","delta":"gcTGH9Bpkz06EhL7X6ObYK6ohJKszYbu+mTgOVRL+yOl0VZ3U5ua9JYH37XwQzAp6j53l4OWd6Ox1CJt4T2lLmJZm/ar8B0+WYxoyqDfXeA/cCnISnR8O34PJUL7bv6OLAcis+PnjU1F7QR5nzd9zff/UlNxJPRVHaxJkCnuJT+jwuxFVRXesttZ8x+NhSgERi48Nz4wJK3/REvo7M/+T3BpqCJVWz1HwTbm4RzA01RnazqYZITBTHQtFwO0T1a/SKusa8zzekM=","epsilon":"A6hcn5MHrmQ1E0+UdWdflBRtGTp+gCvbQ0aBGva8zsAqw0kZEDJSYcchBztVOw/XZxW2iMmt6OuoCnY3s2kNC6F6+LA2k9OFHnXb/nR86yU87Ks+Du3RER142VLsBRB/9U4fuU5f7KItoYGGlFPoG6APZjyXBQsK7XzPqk8meICgb9uxlUPRjlttpdlylz1LJcfnUSrE1kxFfxe4gSKYKOLYFaugfEbpzmp51Jna4ZM=","zeta":"CTYF7Y2DJfkBIjaoiHr5nQ7fTa7jz2OILbT439KCG7BVbskAiBwkzg4MxwTc8NoBJ8GwxMNtWC3vnId2rj0Aokg4jTvHNOghOrt7V4mEh02HUS2gSLDHVZqmNVPeQTW9BO1IWCGPzMDInDfn2UYne8M1q/1wZe+8+COI0xHJ11P24L48ENeWF8t4+bDi4boqkNzUvdwvenw01ipSxcEqjH0UKd2BBlsldLkbFg1fycQ="},"sacc":{"data":"omNNc2dYu6NiTnVYgHBXNuh0hKvcL8wYEwqFsjaiRt3QhAorkWxmMnnaTnugAGwbmDuAOsre0gMKkXw+JXFQm3kWUnAIzcA71NU3yGs8giMPbGu/JLzZN8K/ZKv4Sf8mdFfL1Hh5N/GACk/eCW7uX7w1JwUXzR0+z8ZAp4mrpwSQJhvJRvmuQaRrhPslZUluZGV4AWlFdmVudEhhc2hYIhIg542kR7yDoMEFhYbUEni4Da0kZK3162KvG7Gb79HlbbtjU2lnWEcwRQIhANT3BdUZvFCbrMMRv1m0MTMDwBdeZ1bQeOyDOIIs+jTwAiA9elGZWkmNbUBiCyUGNdBCyeOSltiVLCcWO3lL/I6+Mw==","pk":0}},"a_disclosed":{"1":"/wAAAAAAAAAMY29udGVudHMuYWdlAAAAAAAAAAVmbG9hdAAAAAAAAAAIQEEAAAAAAAA=","2":"/wAAAAAAAAAPY29udGVudHMuZ2VuZGVyAAAAAAAAAAZzdHJpbmcAAAAAAAAABmZlbWFsZQ==","3":"/wAAAAAAAAANY29udGVudHMubmFtZQAAAAAAAAAFYXJyYXkAAAAAAAAAFP9beyJhIjoxLCJiIjoyfSwyLDNd","4":"/wAAAAAAAAAQY29udGVudHMuc3BlY2lhbAAAAAAAAAAEYm9vbAAAAAAAAAABAQ=="}},{"c":"vjAUx4iyASz33XNJiyRwle17UbLoMMH7DIxIWk4FGV4=","A":"PG+mENDoLb+c/dSsA7OBdx4l5FrA/eImXTWHCqaHEWSoCvH5xGacpOIvr47r80q/bx7hrCtFmNo1Q9AsxPHQ1D6i11RkOCLh9kt2jYnHbOymzlBL9XwBhhBu/VcsTe6ead8FnzVEwYcwE7dkPFrGKRUzmMf8/5A+O/TQ1mZlrHQ=","e_response":"XU/sZfwyG6O7Oi08H8aYCD5vQVdbVTjZiLeENcBizm2StaF46axGUpkM1N8e4dvo9PfsLXibIBYL","v_response":"Ak+UvOWcaccWA6CKo/7gAHk942dscSXStCinfUkBiRezeiRXcmrJqWi0IUHaygOiSKPVntsu7TWAkykUGyhcp83JaiDMwj+KincA6WYzd2VuXdDqOBa8AoPQo4e8RY4ZWt/Zxk+geBh2EsDr88KF3KpiKVcmwyWlX8gQZFtYJ3/hOAHlo83b8RxZdlK8WYZem6IWuapp/pFfIpJ660Rw482JO660HaMITAt4d+iY4lFoIGVkOsZSo0xENShRWtSRiT0YZ5EKtdazNp9zFFP2DMVnj9ruCTr+RPyI+lHWIT6u/1QwbTCxB4jwqax3aV/VD8EPsqXCrY8Dj2do5SBk","a_responses":{"0":"P8LzPt8+aQgwuDrOAM51h85UL3vUHS4YHTGAkzP+WvKz8e9VuAuyKvUUvv8NEadG5qFxTjNzDcvBnWVhwqxLjRMbQ0VSncfbJJw=","1":"27rbjbWriB1mQHGq+c/XRsx+HRe86glM9Wokgkkj8HTYhsLq/v9iIBOg1zCPwjwOg8w0keQ0jP2eEr+kNU8ZwomTmsffybdpGX4=","4":"wEz64dp+jYmbB+IOrWYWLueSKrg8y9YVl7kamYOjmL1aWcntPtWF6I7r5c1nRHIIXGIfZ8gVZUXyn0XueegN8PWy8xMmzM9NyCw=","5":"y9I1RUkri0nmw7dYX7Q2cPRBAew7tuHUfEBwMRtP2IDYZ6Fyqq5c3mGCIIqDk53Wmyuj5E7tZqp/kKwW1yKo5zyF4xkHhsbbBo8="},"nonrev_response":"A5k0T8eF1oAc38Cwujk7iAz6Grg/w8KBc/ich3q7QkX05ceSVjhWu7BWhAh5LtXSN0r63YoxXK9ZvJELq4A08v/WCfL2JMc+qg==","nonrev_proof":{"C_r":"BTnSRVm/1NXB7YZcv02+36ulyX6SnWuy/eEx+SwXwwLtTlqcJSue7Bp49M5Ptrza0TkH3GrG5gTiYuN2jN6rGFJbW3ClqnfZNAY1XqaA2Pggxna1WNKd3yAJ4U/HedvnNM7VTnf+JSa6FF/xAet3wW6RAZhOpS5K5UWwHFgMAEw=","C_u":"TZJ0oK4ghkF57CkClTS0Ffv38ykRDI0vgrsqq4UHVsJ3KYqDWKYn/3c8W7K1y91szDCtgEVSFiNarhNpLikkkjagqZFiiuOUYjkt4/Rmo8ypF4qYLS381K3oFunjK8wkPeqxuOauItPuIhQz8zDzwBVzQHwa5iXQz/bm8WvYTNI=","responses":{"beta":"D1691QW1Xddi9/ckbTTlj9ozjZfUulY+J8OHDAW6/OMGKepN4VkfnvCXLI7ljNqTjQhQlBTHi0XFh+bulDIVuW9DQqtao7WLvkuqGzGFxQWIXM+DbhZ9u11d0ceIdSizdG6nQHqS2R+Fx31gQECJcljySzqtKdahhXIPd2C9X45dB+UT0quc8kcOciA1JYidMg2nujfJRl5LWV+YeNWaIZhk1ql++rHrggClzNhmwfkdPmE6n/dURUuh3TJgQFI1LRwoEiVNei4=","delta":"FZJmIFEvnLZXBykPmxaAvxx07YPar584pm+01KshXttE4FcUUVpESuDgNLt7KZZMXk/Rz8/mqLapQQbzX8Zlk6c4iRi5xYCXyYOwL9yu0p9skk3VY55MuzFkgsddW7t39ucGOPDZZo+HygpwDUexxQiOdHxrttHS8efGVt7/SgONOev7XBhkaQmhTqPS90ZNgXWn6GbiDx5MSeMr4FcOvTJPtol48oT2taxzG5MJwRSSxFZoDcAd46aQkPGq78g0Nfn/DNNlTMY=","epsilon":"D9tQ5DMf8QFk5jEuvSe+4IA5V62Tw5z5w+PmOH1Iudk0iUrXqwBCd9vudX8WknESc/8gunPLs661qCAR/1xZkrdZ7Gx6eapX7K8jjD6cRpYsakd7p2Se+sWE9eEPDI4fE7U4K6QBahpK9tpC+O1HWjnW3EcTlWv/vhbuSqS/tWEcDUYXYm2f4z30HzgX7+VCoHC0eyXlFFqEwFbtaTdjG7c9ShK4q6ppPsOX/pSL3q4=","zeta":"FzIxq9MIr22EJGxl3xIzZ/d4aWWKEKP82nebwwQCheuzXn8cANPANRGz/UI90ZS2D1WAsWfflON29KziuV1LoB/uIiVjK6oXQtA8PirprZGdwTe/ODAm9XzFq6kUMNNGyheaBxrlt56mVyK2a8wHB6HHPxwxgIa0FTfDBaIqkinG3K1PDyo5tshZfIVyfvQIjwJwB3dyodej3V/fVeRABjadwexx14Mi+q8puRrdKLU="},"sacc":{"data":"omNNc2dYu6NiTnVYgHBXNuh0hKvcL8wYEwqFsjaiRt3QhAorkWxmMnnaTnugAGwbmDuAOsre0gMKkXw+JXFQm3kWUnAIzcA71NU3yGs8giMPbGu/JLzZN8K/ZKv4Sf8mdFfL1Hh5N/GACk/eCW7uX7w1JwUXzR0+z8ZAp4mrpwSQJhvJRvmuQaRrhPslZUluZGV4AWlFdmVudEhhc2hYIhIg542kR7yDoMEFhYbUEni4Da0kZK3162KvG7Gb79HlbbtjU2lnWEcwRQIhANT3BdUZvFCbrMMRv1m0MTMDwBdeZ1bQeOyDOIIs+jTwAiA9elGZWkmNbUBiCyUGNdBCyeOSltiVLCcWO3lL/I6+Mw==","pk":0}},"a_disclosed":{"2":"/wAAAAAAAAAVY29udGVudHMubGlrZWROdW1iZXJzAAAAAAAAAAVhcnJheQAAAAAAAAAI/1sxLDIsM10=","3":"/wAAAAAAAAANY29udGVudHMubmFtZQAAAAAAAAAGc3RyaW5nAAAAAAAAAAVCZXJ0YQ=="}}]}`)
	byteCombPresentation         = []byte(`[{"contents":{"age":34,"gender":"female","name":[{"a":1,"b":2},2,3],"special":true}},{"contents":{"likedNumbers":[1,2,3],"name":"Berta"}}]`)

	// Revocation...
	byteUpdateRevocation = []byte(`{"sacc":{"data":"omNNc2dYu6NiTnVYgFXDyQLyD407yVkIeGWbI6sih95PpDyKGD+ntn5lhfPbCqHZw9rJTffU9A1pYP5yg8Ps+sQD2HgidJNcLDh32VHRdND9ApB0cbP72qVH9Od5xX5gEPNFOcnbNH3LAKYNI6MgY7mFHQ2wwtpt5U/M6qNFbsR9mU3/lsEfQZCmhQCnZUluZGV4AmlFdmVudEhhc2hYIhIg971Iur8V6C3+2XB75zxPfFZvfEtnnVrlgonZsp/76a5jU2lnWEcwRQIhAPw/Ln8rk/zVt3ZTB0yI156LlrH//NYpGOlTymJYwi/kAiAt60syH4HOEPbXJ7aZLnesQe3DroWamRWsn7icvFaxMA==","pk":0},"i":2,"hash":"EiDnjaRHvIOgwQWFhtQSeLgNrSRkrfXrYq8bsZvv0eVtuw==","e":["BO2VVOUEeWypMB0s67YSKi4K1nTEIkCmpQ=="]}`)
)