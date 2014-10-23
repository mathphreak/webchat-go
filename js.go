package main

var js = `H4sIAAAJbogC/6x8a3PbttPv+8zkO0j853iICqbtnleHNKOJHadJm9vTpLcoagYkIYkVRcok5at0PvvZXQAkJNFJ+pxnprVIEJfFYi+/XQC5EmXvk1zWRThZ5XGdFrnL7s1jb+oKdl/KelXmvXyVZWEohh/qMs3xg/929LsXiyyD5/F67RTRPzKunU3Teta2dkyhE4bYtq10tzNEPxQHBwIG8q7TPCmu25rXnTW9vEjkx9ulxCbP353/9ubi7ccvb989v2hbXlqEaCp3yXhmdY7VDw76d/Tzjup7U1m/L4u6qGGkdxP4Eob6y9IUt52dW+Plq0UkSxgPqxSTnvAymU/rWVu7tIZONEN5sxp7c95YVC+tz8Ir5TITsXSPfP9oyp0jh7VF7ujZ4afxgKnfkTi8GzOo9PiR8+Tky5Mft6vC18/JWNfFavuVvmDpIZTUxeviWpbnopKuRdnKpqyX5r13w3cjMfbxT5jL696vcnpxs3Qd9+/1588VcwZi4LjwtH7CHKuf58CLaJeZfcPMaL3+Y4RMGA8jPxo4yxtL/F5YqxDP0iwpZe4AJWIYay57ppj50luIpSl5CyJVbS1BOnFPUMCMsDEzNXs1LpBWPgH9KUo3w0lHbHJw4D5zo1E2Zuv1T+qBDU0RSBkInnqi3/B+w/hPzceftj+Oxoxf0CPHGjAU8/GhH4aVqUNNW5pubP5pHY6Gjx9JVGD8403SrJalG1mNrtREeNw0nMH3YdQIJ3zxo7b+rRoEbckEmBRnoqreioUEo+DwOAQeTLwIBOR3kSGpAXAzCuHBcDEeNt/9SWC9hZFvdRdaQ77EVcEBo6AubxtJGzp1uZKgbWK9dp2JyCp6GfZPfAcnTy/44B/9fXzk1bKqoaP1OoUB3gJRb0nAoIgNj/4efR59vh83tYbSW4qykj9/ePcWuSf8CP5sYlHHM3fSSvumpfKVYkwE1QMUCsWhY+CJLWnaIgST0ziYDAbs1ZYcjiaw2myDTSueccnPeA2SAH3UXpWlseQJPKlV5PNQWU0vKeLVQuY1X4BI8Xf454/w8aN7Jy6y1SI/jItVXjv+CVfvFTw5kyKvD69lOp3RFydLc3k4a96LpYjT+ha/3B3CKPIGS++KYuGfbPj78Ojvz9UPp+7n68G6z0Z/Px3/8PSI/wzFVMbg4+ej4VN36J9+Pvp88nTNnhzxP8OjU3fYF6UU66hcAzFrCSuQrGflOl1M12m+XNVrIGS+XsharGEBxIK57ujztQ/WjEZhn4+eHqVT/isMBb1HRXK7ntWLDPpP+V/hUWvGPoXOlch6cVX1sEKvljd1LxG16F2nST3rqbn2ismkkrXjVcssrV2n5zD+X+Hci4HIWl5kEvnqOrWIMgmfPnR8KqH8Y3hfl/7+N6QPPtOv/1+8nkmR4O+kKGr8TfwPUAh/nB+c/eZJegXmkQsRHsXFYpnJWq6zQiTAsjQHCRAgeFfyiEdQAQT4+nD8A7D58aO3KAG/h2/BXisHznMsec3f8N/C7lF4LGAKIqKl9nG+r2jROdRNijy79R18egdPDgpP6fgOsvUFPHGHtBZKGu11+ELcKEn3HXh8TY9gH2SWVSha+RRqw8sH9aK+LEWSNF/eqxegoLiGJjkQUFxD9dxBMVYl8KBKVpUEk+478PtGLB0+AcmRUVEmsvQdejmjF2wLnIP/kpTWFPugggtd4Gz4T+GzshS3XlrR73rd4Z/RzVW1yGP0TFRtE+TgVsA8yKqFV2QRwAL2wXn1wUyd9DscS/8kMNb0WkbztH6juvkgM8AcRbleC29R3HWUFntljx8JQ4UpRAs8MdZ3Ysx6xGjQOEDzBJoGLEALFLiTsB+jE4rD35gnlkuZJ+don9CsxeH/zb3LSrgxtPdIVggfBWD1fwO8sCiuZFPZ2PtN8Drs4mCLLw4Hrgc2eGp5Yd7a2Am4iLr4DSjRwMN3HPDEwZuuXvcxVWR1JRqaI8B0E+wl90A8pqgJW6uGjZA/v/CET4OfjVsAvvwSSndXhxS+8Z6cMMaCX8AbNZNDTx22U/2TO6dPTp6eHj358anDGCfPCHWi8H07RNsbjxBYfIQOoxAsBOPT8OMoGvMpzCOX5cuPb16HDqAp0Hkgy5MCnJPGO1PLqzBu4f3p1jrVs7RiG6D6GcgIEJLA7H5hXPc14XuSfHhy+slaeZYAxgNuBhK8L3Bf1HVJdbFPzfVfkMufdpRCoNcejQPhfflCqPrLF5jBJIeCSgtuGCGmCBpXC72k1afwW9oIY1HVPK13xiQtQ3UQDP44FdlGC61jKa4WlC9cBkuFCGJ0PD44aFaHTcJWYqDXZq2gfy5CRBvECmRV1LdgjwQmgePOSTWwE1SjOYewxFTHCKqtPUf8LZJbrA4faepWBYGFPymCHj9CEKNWAEovVSmsi00QfnloFmbC/wOTodn0DAL1PrmNFODLBKsEcmdZmgq4ZEp4AhDAGzDLydZqEwTkk9BA+nJKsKfiJyxwoqLIpMit5SS9EjCRiVfN0kntgkhOPPBdFyTbpmPQ9AvCuSi0W9IGs+qSIC5D5z8wTgSiAXisD0rueOZ9HgISTgA+E1xzTwA7g9ZGQrF+3gxwTao+PTiQQxctP8Sc2pqc3b5KoCIbAhj0R2N/22ccHPyfrXeA91DHKD2EEXKYDO3uqrPbc+OWoVt/59tHMaUvCG+9y5Usb43feAYdRozWAn2kABUD8GCwpu6h+TTsXFMBtEYYuJtqGHr421URKwcRLVZkeSIWMB06COMnjwPjMJEoXOZwCg9p9UJ3GM7o9Q+CxeEdvZBzDn+i5/cZ0KBC+fAZlVwslvWtLrFFzQB4MsCiddSGEiQgzVXfu35Dz702ZlKtTcTRqRE3gd8ZurLwNU4DdC98IAOAAYzj+ArEtckHamOyMZsAAjxvtUqT8Bi6q1bLZVHWgPdIh5alegJ01CnKMUYXCRqTc7ILMG3sJzk1qYsggSBlEkJUM0qgJuMqMwEOP/aWqwqDIWVfqKni1jeq6wken8Z6jCEafhQRCK8QcmS3LsY8zI9pEqCtDxnyhmgItCCkMkSDGqDWmBgWoyowPfDDKOZqCZ4oghGhQcjf1YBFYIbnllmQ3hS8ebcJS1p0tQl0cIYRJAi3FU+G+Ee9Mj0/1zlT5qungtKeWvLeiyZlhZLWew6YQydTekpoexdlCRDcDl8eP9qm7e3IGalEWM8ZRANn7IxhotupHKRjAkGCNo5+bcwk4P9kFUsoUA8cFxHe8IdXIGjwjD9cy7rfSD1X6+nXemE5AvQOOZeuSscgEnkItWkFUiYaIiKyrb4Fapq+Yi0/1FvjILANedMtAoQxysrTfqiBu2CT5x5Ga0PhSgYhGUQhF1fQx+u0An8kS9d5/u7NuYoZXlMs5tjoCltteP+kkXIkZMPB5HbNXSD8MxkqgmI+/h0dn4LiC18M8E3L9HjD64LEoGPiVA8GgaWE9a/SO/lQJZ2Q5AoEPlRLbnlIdq+0mD61FvrgYKegC1luOPa1Nfnak1fgZ9pZdy57X/m9raXfZqpKg3TxFaHUkIjLi9ptHsBW+NJNOge2YIiOn9BiK3EDGfA79V26b1TvSsRdSV8ZtkmrLsKOT61FODhox1LLPsYROZDq78EesNOBBonkT0VG6TXWsVzAQmHNEewBW68jZX71qihYp6LOfRysWKdzhYL55zQoDO2ltVywJqWqkoqBqr0Lq4Dw46cTK1Q4ONA0IFc3FpLc8Jno5JZNxr6gXqo0XYMslMojTTRJg0891AZXSeJlxyiHaPj1lBVmE0oLzRsfiMEJQ3Erq9pWGGSfCPXKNR4CfA9RJog5Gw7Y66FWliwcnjzcA06kQyCwddDmQpu9jmYVVeK3g4EtCY1CgRFfyEbZ9rXC4rNSiw3zgXHWDIbSVXFAK8qaj2jd95evrUz9AdLd8DgrKmmzy0o1m27jEGAY+u79GVPWAudMCYm+Gw/BkMaNCE6Y3yochiIAMSfhREHU/jUFwBM7F9LIKOA2rsq35bRBiYiiFInB8enEQCfoXfm2ye7+AoWZFtal8VFFjp9GdnStKWgUh4vWCN6oldCEdenPjTFPy2wVz12nHRBzD48fYWuzK9Ld/ivr98LYEiJCZ9GqhxxKdx+W57PTFairVRplYJWq76XLzhhpC//CtVeTdW2zYXyCaSA1CYnBwPe6RCq00zBINTG6E+V0YZzmawTBOk16Vlx/7/hOXuS436FMVX2bSS9JAQaKW1CD/UIgkBIJptnjRwAXzovFclXL5APWVNRBNb3/uZRlffu7yFYSc8TUCYiNZTzUAFxZIgpIMZLEDaNgMYrG6zWg8t1MWcS4wldbqcUJA7Xeo2fyLWL45CHwAR2aicZkFZwoK+I5NEHKwpgFomEMlmw2hBApSfdHWs8e9EaRBJ0ns6qGQ8B1Xe4gW2USZjp1ow0XqHjElCmj9ojWjnHe9iTW6xMbIQQPC0DC7o2jw/EpWB/abj+B8GkIsWcGbMCu3f4x8yekW48f6Sb+jlnSpJqu8bGZsTGsxuQFmFxpNlVdxhqrJyC8IGfpAs7Aqaql1vZiC8IhHa9Qhx7i38MM0AlaNPjUMU8wL6TNkItpqW1+gPUXQWKcVdJwDcpjQ6FizyrfXlGt62qVXLaniWYdLPlxTZnFn134Sivz7+ylMpT2iiK9szR5sJ+4qlqFwRMApBUMY4npNJMPyvnuHNWS6FkFrgpcjIpF28MAo3H/GI2Zy/zIQ/oUAFuW8qo7BtxyU1ArLVaVthkflCcA29NA0vUas+GAlOVN/e3+sNa3+8LdrE7ITikCE0kOO7izBVFalxAYGSC34jZqcKVBDyWLEPnQCFg21ErX9uFj4AWL9cA8v4u2CGy9NSZVwg51EKvxu1USmpMDjuM7ziDqJNGqr4nE1H8XctNSb8KM5jwHbod9k/6JOY7ROhnaOWvy3Ob0hTDmIAPNx8MRgZrFRPWqrPUzIDGNwMO4scH4srYKMUFHuaE4BDwRaq5h9GkC7LYyDM7/m0PoDyYKNLNDtjZp3b7bQN69cSEANOsBn82S4ImbiV+Z0P7Z9nJ8Xb2RhMePtliso/vdWQmtyMXS393QiQWQsF4LY+a+a3WJeDHe5TUUGT6R72znuOG4id8ZIjj45dAZtPttf3Hn8MnJ7sklWNou2mjvChMNvvXa7l9QCgQW+yW6DGDzlfjv2wsqvUI4E+6qJpVuzd34Ym+xyup0mcmh5Z4pynWKJZ23Yw9Hy1oUMbUvkw0ztpFGA4utFh8H0eNzdTBiFyGYVHjvYTvYeOQgsicXe6pH3IXBauiuVcl741hNBWh4X4OAYZZyeZjgX57JSQ0F+HOY0M8mAJMi6jRGkKc80LKoUsUI3E7wzGvolDITeFYCMLCqStkkTPtamsislQ7aGFtr4FmxyvFAwnmWArW/Ah+bva17i7iBTvsuxVT++Y4mxJu52B//0h/pNIr/RtQzr8QhgG4qYlydUNn+pMrYBuBDVfl7u7M/Pm0kz0xKr4kJn5N9nB0jLlDRQT/WTAi6N0ibuI2g/+i1Og2a7IN0BX1pg5IImOsNCRQYsbOrPCdI3nQZPdClnQCcb+iM1AKimm1K8aAni9brY8zmDxchnhccgBsbqLOF/sOhnIqWlLkzA7vY3KTLepavQQejBsGn4WIAA2VmoIzTQb6BEzj/frxsFyh+rTXIwEfww4MQhhos0DCr8ztdke5QoxOVW8AVQvuhFb8Fty1gbeoaW0N5OtpLfGAAK4m0HenueB/Kut8Sa/mKsrL9E8qyfq3zTnuDa98eOtI+kd2fYbZUyf2ttkW2JWJ6u+ToczU4mrLO3KUBb2bKdHLwrElgctXrWePE9ZZDPHDjodND4MQGZ94/RZrTrgxorBI+454fnOq/migdWLA25m+bID4465r62fdM/QygwFnjQnGBOE4BxFJ3f6YPDDAKm1UksTsfexO4a/FiZReSxlk0VMZcE/49tKqzpcN+Yi+UHzEI8Iw8YUHiWVx3df5SLUcVl0WWfSyWXcHwlhGNQqep7LT4K9jazIkahNx2bEoss28Zhmi4p9ymZSg2/kMfXevtT5WSNZN5je7ou2eDte3pPH701QlR5/aM/vzuGWHTr05JcOv9LzUl48TtVl3zaTLxJiG1DS54bJe6mC74VRmiyOBuVC02JORxrPDG8Qb8doNKYoIjIW3cvsgKtd/DFPJYgN9N80OoALoOniGIFVT5am2soasTyBls144I4FEDdZQRuz8kcNC0IsjxrWZEid1OwxacaWxhrMePaNbxFsraGCj4fifP/NWEhUkTif21WK9VBjDAjQ69BqJdgIODFta1HGuRHeaXxFaHgXU+f0Mb2V4iazw1QM9K84ORozjAHX3KedxpUJRqWIcTvaPOxDHFwPapRL2HjrHMnpHj891toph0686ds+F85FC8D5H22L+mkt1TNiOtqlTFTbYlGTiGh/D8bscxbwzsnDgpuDa0MbiDOR7ew+1KtOSg+CNHTGo8JYs5GExWwJPK/8GDSl90so030D/6Xz/us0ErKKicvnnRHNyyKRUA3iIDfdrNHUzQCNz2Vs8mNxH51gk2hIZ8wbOwM3t68rRJ++HXLkYJPmX3izAeTv2pvQM0DY9xsKmHOSSdPMJ9LyyiFCelm/0fqcRvAoc8tPbL9laTLxhobpfwgV3LSLy3cnzN+b3+oj2AZ2WggwVAtQrw45lO1cJkgvzg4NXO1SJ9oqhRNWDsh/NfX73/6PQx+m8/bAs2HjSuKSHgYNLn6B9xJUAc02Xt0A49fsM6VRmv1zrAkRBFKiCoCrhoc1rkfDdMaeooHoqB87EAwKTm4AzAfThqJlD4jOSRWcIUITaLGEqukms7qQpSnHuf2pNS6khp7q3y9DJ8E6BhgEFSkaV3kqKK8CWQcUd34vJmv28DPNUTUdfl6G9wVaRJDzG//vYEgkzzqKqgDrULOrUu183c2Nr1+nKXJuu1qx7C1WCwdUcu5gmXKJBJeO0mFEglXl7R5kG2fZ9q6P+97jFngN8bkwVgDQCb90NvCBgU6/TonpWenns1QmLGePp2P19gUQmzA1gFSxt7MgzhiamSvFqvM2W2YySLiqEaHj2e5Az4M3OlKp1i40qC1E2ZfXnqGseJQ9dxBg0Yd7yGxHvpx+jI8wp8kTlHSceMICppIPXGvvPX0l2NYpjbSxB/eojta3+Ks2Ao+JKvlFX6OURu8Dq8Gv2MW1b4g3evQEMb9HnUgT0Thcnp5BBtpyVGN6eu0XdzhFeqjMdMLecMT1pJ+EHGZPArMabEU+By33NAXEy5C5l8BKspawrSQXaSvsoZggud7hyBSJptYxyp+zzURlH0c+jOQCWycAlajiSB4txYBxoTlMFSyyAsfVq9WixkkgI9GLKKqcBqHwA7LGUCjuQ+8TAFFy6odxn+rEcHvntfYPQqDJ8PR8nYh//NmRn9hbGAjuDhifnEwy0AIPS5nIhVRmgNIt5iaQ3qtufLJXj9mZeGtbH76niaO2OBs3t8C4FujIcR9851XQIrJOOaB5z4AtwFQdIr9Gq97vdXW3J8TrkT0laeqRVbkjwFboJn19k3ZajV9s7PQIOsZe9qtByPgP3jwFFGv3tOHd9gWqi5PNHTStppJda0MmWSrWuqQBZJOMyjH8PC66V4rxZGJkx9AzAC2jNV7rRu3c0c7ZcS+mQ0Hwcx/Al3cP9oOgb5N7KKfFbSktiCCi2hmtwwIn2bhj6K034xXlXCPsnAI3OSIYkWBoFNua8zTLacIU/2CkHQuuYPGp+1V1ysS7oN5pPhfVGmgPNFRkuCp1mbQ7IxO1MmNAHNi0EdgL7nwE6Jj/jeSHfpSh6rK4Gr8IQ/5y/09aT2OLC6GXgRTq2j0Pwm3Dfqexm9eMOv8GrYLf55GTpFPgHbVaV0d1a77ld4LBQKfcd841G2KvVrsaodiDvD+0WxqiSyBr7QM4giIEl6zKS4kqYYGwS3AHHSeB7eelQI4+TmebU0TyjNofMGH4mDlRNMPeJ/eI9n8p6ZU4znm2C6a7tIgJUZ+nErJX9w8GL3/sCPZOAugFGqySLsOqtipPPxI6nOnaH52u2LMd+W3oVy8WQTdEeLDYx147ZmWg5dCQhF3VHAteekUjQfPShCNw5O1dfFplpQz8riuoeYAPds6DCw68ibJaX5e81FfBYgh8Ab4AW/LR5Br1txXFOqG6zyvSY7DSYTKtT1Abfu9I+WoGsEsJkn0Eqhmj1+9483XHYUn2z4Gd79pHuea0uhn6wzcSvL0Z9/jZ+wI3Do99tOBADmnhI7HL1Kl0/D2l/xdbiVvuOQsEVHRc0WNOTwQX6L9wnEZRbvv68ZrMBWuyy9kjsr1oISir1ZQ5BrbQwVJcfKW4hak59+s8OW2O/rssi3JzUHXKZTORm/5qU6TQeqAq6tf4OqafCVvufGrebIEa/pSEJH4DDK4IbM64W7gL/kAhZgYhfhHOLxBKLR51rt5+bznKmPzxUWWQAWWYSyNcU7UaMEKSjZPbguN9uytucuTJsCI4z1jN4/CMUSDCT2sJ8EQqaAqWsCfszThyNhqUtzegjJnx8czMERlo01ATeg7k656Iz4fbwqMZ5V+NGfc5SOF2kJfrIEJrnZer1gmrQ5H8nxQ5aNn1CmIHiGk4PFmgOvryEUoFhOLelksrOmej3/7VpmxqqgTDOeNQvZrhQsZLOIqgSYOLcWK9vbODk3iY85mitDNHjE6VSWO+Idh0jmej3dvjUEZRD5kzPC/X28eRdjYlCB2+Sr+zZ0PgZPn1Jzxz5Q4W19wp6nevtBU/dS5EkmFSrbpVx/6/J8wLkOUmYgQPcyXLo3O9OJMYVlpiK14IVloFfprpFqgCyMb4+XhRpiYrQDKy2/GSw0Bt3a8Mo2gUEZPQMvenjlvVdKPMfdU3mw3iqnQsIQvSTK1EMDJXoaSPQaGNFrIEnPoJBei1h6LUrpxTORT2EgMl+9ubylDuEX3ElV4QP0K3dv3eyj9xiv2mKCYxxaN3Ja55molUfvShxUaSm9ptAacxkjxQzMwgHi6kq9WaNsxQ2dne/tJJa3CojHY5fpf9sCGm82e0kVLSQ7MqZUBEMAHsPyo2QwHfk1zkEf/yQRu6WI3NFYDm8z948pu8FMupi2VxPmRKsoymQFQDUbQq1+MsrGvsR/8gSf6O5dWmu5BRyBSbIWLiPr3O58jNNcNgb9u9+gfmtbOVVZIH7/ybcU5RKseziji8pNzRlHjjMVwKuLypd0UXnmffnyKWxuKfZmeAvE7q6B4SafCaYamQZfDg6AtE9oFGbIb1yZvX16lX8yC6V95d3efr5JXe2WhzYjrH9H5Q7e2z6tMxCbzRYjn7jd/2rU1LjtVXiNWZvLcGXtYDUvLynrfrhC4A/rpl6DJ65zWpdPT+vkqTOYDpzTI3hSbzPzdgQVHHN87mPhXrPgEoy9PcpXB1HB0yx0RFoBeyLAKVm6WgTxnSiD5J80z4OL1XJWgpxWEKHN8mCaizK7DdJyHsyKVRn8IzIBgxdVMM/B0BzOcxHPgywTCxEscjAwUBoA0xdFUKzSf0QvKkSZBMtcrvCbCC5FHYA1xdX5cFcGdbXKxSIN8KDFVQlkiUWU4fX14BpcdfBnKvLgNl3KMq1l8AmAPP6Pu47G4gROk6EEb43/ukeghAH4+Z9FNQXluqbnrIBn5PJ/ctw4Zni8x52N6EzJJCtA7dTxEoHUu+yHmdlTGytNvgzxfjF1ADq6gA5g+Z1qFS3S2rr4NlX/ygXSAg76jkZh2/+8xSzUxSAgO/SAB4WGoE6gXXQ3UoWp6eTWvUf9gMgRTK+Y4j8rUk39Gce2/iX+EwOqT9zvnu4ljdQJH82mP2T0oYjnEoJ862KmEtxn/NyONIijGFI1bVznuvKPjpxBVsTkxbxZUdUD5+i6cjTxRV6AhIS2sk9BtqOneK1wVZ0eRU/BjMP7OdSWSncAzYkoS6uZTPpUgQWVrD+mC3RP7g5FX2MP5kQdxZVd1tIFthOwjzG4t9L0/QwTK4psgpXfoPvxox3KqVHiaaKfhRbZ5/zk4n+33euV2xpA5cnaW7CYhBS1AJOKknGK/7bLEWo8bje4Sw+nwzw8n+uqKwxGGsCALpXXWXY2BGEx7fgU7Pay3UpRLGs7wFtgl82JF7THLnISUEFZNbYHhshSY5nwCWTZpBr5/98aWZQRDLFIu6fJNaSo+bTJ/h2S6OzHXv2l0ubV1hyDw5N+uIKZXkLx7kYAWXXjyPpgBVYoSc2yEgDaUZqqyKQHH6xkCDQ4bzdTilxtJir0Fjal+v1fCLzi0cMSr07xdivgX2DRexH+q0IA/JJCVr28qHv63nyr81UPbOJRUfaIAmxueyPaLv1/AQAA//9SKftXUFEAAA==`