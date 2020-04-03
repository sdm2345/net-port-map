s = '''

00000000: 1927 d040 1784 704a 2ef4 67a7 bf45 5ee9
00000010: f0b3 9cdb 90a2 37b3 fa3b 35e5 7d6a 3b28
00000020: 3622 fa28 237f e9fa f007 0b9b e36a e8f0
00000030: f877 5de7 52ea 314f 6745 84a6 1bd1 92bc
00000040: 3cb0 b294 e22a b5d7 1ba0 958d 7ea9 aae2
00000050: 9d7f d2e4 1cb8 2e8f 3904 2c37 3934 25a0
00000060: 5fb8 6c92 edf0 51d1 5f4e 6228 2809 d463
00000070: 0e86 edd0 2c9d 5dd7 92d7 2493 e178 ce7c
00000080: 6ad8 d02d 55ff 80e0 9865 db20 93cf 5e71
00000090: 5294 fb46 e588 86ce 7cef aaa3 1fc6 8a72
000000a0: 7405 dc8b c17c 0925 394c 84d3 80e4 bb31

00000000: b1b7 a46e 8fd9 12ef d105 b985 cecf b006
00000010: ab71 9792 96e5 9b2a 8477 f8fc 8cfd f820
00000020: b5ef 35fa 5b76 23f7 6bb6 888f 025d bb9b
00000030: 8e41 db23 dd4f aff3 ef43 c637 8130 5dff
00000040: 9572 eaeb 2f53 f39e 58f3 3503 f42e 727e
00000050: ca99 dc64 0346 278d 5b80 deee c06d e0b1
00000060: dc06 b17a ee33 0b65 340a 8436 24ec b824
00000070: 75d1 e74d 0591 ff76 93c1 975e d729 4a6b
00000080: 4f4d 3da1 f2a6 fa17 c0eb 2bc8 a9c5 b240
00000090: 54f9 2803 04ee cfaa cfd8 6850 6401 8a87
000000a0: 07ac dcfb fd55 aae8 68ea 63cb 9ffa 626c




00000000: fdf8 a38a e602 dce1 02aa ca9a 38e9 8e18
00000010: 5898 bda9 dda3 55c1 61ec eb89 a7b3 9897
00000020: b90c 008a 93b2 da1d 3799 a286 666f e91f
00000030: 5b50 49f0 d75d b889 9db6 95ec e0e7 121c
00000040: dcac 7e13 0cdc 71a9 471d 8481 aedc 58e0
00000050: c335 3f53 1ebb 916f 6074 17eb e873 159d
00000060: 1bd7 20ce 1f51 6bcf dff0 0211 f22e 62ae
00000070: 091d b538 e15c 80aa f5b2 c4a9 d24c 3471
00000080: 8738 e86f c6af c0d3 b4c7 484a 15bf b884
00000090: 2144 392b e7dc e8b3 31b7 e9b1 0b87 e38b
000000a0: 3737 3396 3e5e 87eb f598 7ef8 de6f 68d6
000000b0: 61bb abdb 23ed f8bf fc2d 490e 9cdb cb4e
000000c0: 717a d957 fcec 165a fb7f 64e1 8477 dfd9
000000d0: b8d7 cc61 8216 b7e5 0885 1a92 c889 c5a5

'''

data = {}
for line in s.split('\n'):
    line = line.strip()
    if len(line) == 0:
        continue
    k, v = line.split(':')
    if k not in data:
        data[k] = []
    data[k].append(v.strip().split(' '))


def start_cmp(arr, line):
    k = len(arr)
    for j in range(k):
        for k in range(k):
            if j == k:
                continue
            for i in range(len(arr[0])):
                if arr[j][i][0:2] == arr[k][i][0:2] or arr[j][i][2:] == arr[k][i][2:]:
                    print("cmp ok", arr[j][i])
                    print(line, ' '.join(arr[j]))
                    print(line, ' '.join(arr[k]))
    print('cmp done')
    pass


for k, v in data.items():
    start_cmp(v, k)
