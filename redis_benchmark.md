#REDIS BENCHMARK

## 10 bytes
❯ redis-benchmark -d 10 -t set,get
====== SET ======
  100000 requests completed in 1.28 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.183 milliseconds (cumulative count 1)
50.000% <= 0.335 milliseconds (cumulative count 55536)
75.000% <= 0.359 milliseconds (cumulative count 80151)
87.500% <= 0.383 milliseconds (cumulative count 89081)
93.750% <= 0.431 milliseconds (cumulative count 93754)
96.875% <= 0.527 milliseconds (cumulative count 96939)
98.438% <= 0.639 milliseconds (cumulative count 98472)
99.219% <= 0.887 milliseconds (cumulative count 99220)
99.609% <= 1.199 milliseconds (cumulative count 99619)
99.805% <= 1.231 milliseconds (cumulative count 99838)
99.902% <= 1.247 milliseconds (cumulative count 99907)
99.951% <= 2.055 milliseconds (cumulative count 99952)
99.976% <= 2.415 milliseconds (cumulative count 99976)
99.988% <= 2.679 milliseconds (cumulative count 99988)
99.994% <= 2.839 milliseconds (cumulative count 99994)
99.997% <= 2.879 milliseconds (cumulative count 99997)
99.998% <= 2.927 milliseconds (cumulative count 99999)
99.999% <= 2.999 milliseconds (cumulative count 100000)
100.000% <= 2.999 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.006% <= 0.207 milliseconds (cumulative count 6)
9.576% <= 0.303 milliseconds (cumulative count 9576)
92.210% <= 0.407 milliseconds (cumulative count 92210)
96.346% <= 0.503 milliseconds (cumulative count 96346)
98.136% <= 0.607 milliseconds (cumulative count 98136)
98.857% <= 0.703 milliseconds (cumulative count 98857)
99.121% <= 0.807 milliseconds (cumulative count 99121)
99.229% <= 0.903 milliseconds (cumulative count 99229)
99.263% <= 1.007 milliseconds (cumulative count 99263)
99.292% <= 1.103 milliseconds (cumulative count 99292)
99.681% <= 1.207 milliseconds (cumulative count 99681)
99.950% <= 1.303 milliseconds (cumulative count 99950)
99.951% <= 1.807 milliseconds (cumulative count 99951)
99.955% <= 2.103 milliseconds (cumulative count 99955)
100.000% <= 3.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 77881.62 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.352     0.176     0.335     0.463     0.759     2.999
====== GET ======
  100000 requests completed in 1.38 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.159 milliseconds (cumulative count 1)
50.000% <= 0.351 milliseconds (cumulative count 56796)
75.000% <= 0.391 milliseconds (cumulative count 75586)
87.500% <= 0.487 milliseconds (cumulative count 87796)
93.750% <= 0.607 milliseconds (cumulative count 94010)
96.875% <= 0.735 milliseconds (cumulative count 96907)
98.438% <= 0.887 milliseconds (cumulative count 98491)
99.219% <= 1.039 milliseconds (cumulative count 99222)
99.609% <= 1.191 milliseconds (cumulative count 99614)
99.805% <= 1.335 milliseconds (cumulative count 99815)
99.902% <= 1.447 milliseconds (cumulative count 99908)
99.951% <= 1.535 milliseconds (cumulative count 99956)
99.976% <= 1.615 milliseconds (cumulative count 99976)
99.988% <= 1.687 milliseconds (cumulative count 99989)
99.994% <= 1.743 milliseconds (cumulative count 99994)
99.997% <= 1.823 milliseconds (cumulative count 99997)
99.998% <= 1.855 milliseconds (cumulative count 99999)
99.999% <= 1.871 milliseconds (cumulative count 100000)
100.000% <= 1.871 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.013% <= 0.207 milliseconds (cumulative count 13)
6.358% <= 0.303 milliseconds (cumulative count 6358)
78.764% <= 0.407 milliseconds (cumulative count 78764)
88.976% <= 0.503 milliseconds (cumulative count 88976)
94.010% <= 0.607 milliseconds (cumulative count 94010)
96.375% <= 0.703 milliseconds (cumulative count 96375)
97.811% <= 0.807 milliseconds (cumulative count 97811)
98.598% <= 0.903 milliseconds (cumulative count 98598)
99.102% <= 1.007 milliseconds (cumulative count 99102)
99.429% <= 1.103 milliseconds (cumulative count 99429)
99.640% <= 1.207 milliseconds (cumulative count 99640)
99.773% <= 1.303 milliseconds (cumulative count 99773)
99.880% <= 1.407 milliseconds (cumulative count 99880)
99.940% <= 1.503 milliseconds (cumulative count 99940)
99.975% <= 1.607 milliseconds (cumulative count 99975)
99.991% <= 1.703 milliseconds (cumulative count 99991)
99.996% <= 1.807 milliseconds (cumulative count 99996)
100.000% <= 1.903 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 72358.90 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.388     0.152     0.351     0.647     0.983     1.871

## 20 bytes
❯ redis-benchmark -d 20 -t set,get
====== SET ======
  100000 requests completed in 1.26 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.151 milliseconds (cumulative count 2)
50.000% <= 0.335 milliseconds (cumulative count 61535)
75.000% <= 0.351 milliseconds (cumulative count 78624)
87.500% <= 0.375 milliseconds (cumulative count 89917)
93.750% <= 0.407 milliseconds (cumulative count 94073)
96.875% <= 0.495 milliseconds (cumulative count 96985)
98.438% <= 0.663 milliseconds (cumulative count 98454)
99.219% <= 0.871 milliseconds (cumulative count 99254)
99.609% <= 1.055 milliseconds (cumulative count 99613)
99.805% <= 1.255 milliseconds (cumulative count 99808)
99.902% <= 1.351 milliseconds (cumulative count 99905)
99.951% <= 1.415 milliseconds (cumulative count 99952)
99.976% <= 1.631 milliseconds (cumulative count 99976)
99.988% <= 1.951 milliseconds (cumulative count 99988)
99.994% <= 2.023 milliseconds (cumulative count 99994)
99.997% <= 2.055 milliseconds (cumulative count 99997)
99.998% <= 2.079 milliseconds (cumulative count 99999)
99.999% <= 2.087 milliseconds (cumulative count 100000)
100.000% <= 2.087 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.008% <= 0.207 milliseconds (cumulative count 8)
12.939% <= 0.303 milliseconds (cumulative count 12939)
94.073% <= 0.407 milliseconds (cumulative count 94073)
97.155% <= 0.503 milliseconds (cumulative count 97155)
98.187% <= 0.607 milliseconds (cumulative count 98187)
98.634% <= 0.703 milliseconds (cumulative count 98634)
99.032% <= 0.807 milliseconds (cumulative count 99032)
99.393% <= 0.903 milliseconds (cumulative count 99393)
99.566% <= 1.007 milliseconds (cumulative count 99566)
99.659% <= 1.103 milliseconds (cumulative count 99659)
99.765% <= 1.207 milliseconds (cumulative count 99765)
99.861% <= 1.303 milliseconds (cumulative count 99861)
99.942% <= 1.407 milliseconds (cumulative count 99942)
99.967% <= 1.503 milliseconds (cumulative count 99967)
99.973% <= 1.607 milliseconds (cumulative count 99973)
99.979% <= 1.703 milliseconds (cumulative count 99979)
99.984% <= 1.903 milliseconds (cumulative count 99984)
99.992% <= 2.007 milliseconds (cumulative count 99992)
100.000% <= 2.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 79554.50 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.345     0.144     0.335     0.431     0.799     2.087
====== GET ======
  100000 requests completed in 1.26 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.143 milliseconds (cumulative count 2)
50.000% <= 0.335 milliseconds (cumulative count 59748)
75.000% <= 0.351 milliseconds (cumulative count 79807)
87.500% <= 0.367 milliseconds (cumulative count 89770)
93.750% <= 0.383 milliseconds (cumulative count 94090)
96.875% <= 0.415 milliseconds (cumulative count 96965)
98.438% <= 0.479 milliseconds (cumulative count 98439)
99.219% <= 0.623 milliseconds (cumulative count 99249)
99.609% <= 0.711 milliseconds (cumulative count 99620)
99.805% <= 0.847 milliseconds (cumulative count 99806)
99.902% <= 0.959 milliseconds (cumulative count 99904)
99.951% <= 1.271 milliseconds (cumulative count 99954)
99.976% <= 1.367 milliseconds (cumulative count 99976)
99.988% <= 1.567 milliseconds (cumulative count 99988)
99.994% <= 1.695 milliseconds (cumulative count 99994)
99.997% <= 2.007 milliseconds (cumulative count 99997)
99.998% <= 2.031 milliseconds (cumulative count 99999)
99.999% <= 2.055 milliseconds (cumulative count 100000)
100.000% <= 2.055 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.013% <= 0.207 milliseconds (cumulative count 13)
10.781% <= 0.303 milliseconds (cumulative count 10781)
96.565% <= 0.407 milliseconds (cumulative count 96565)
98.702% <= 0.503 milliseconds (cumulative count 98702)
99.166% <= 0.607 milliseconds (cumulative count 99166)
99.606% <= 0.703 milliseconds (cumulative count 99606)
99.764% <= 0.807 milliseconds (cumulative count 99764)
99.859% <= 0.903 milliseconds (cumulative count 99859)
99.913% <= 1.007 milliseconds (cumulative count 99913)
99.922% <= 1.103 milliseconds (cumulative count 99922)
99.927% <= 1.207 milliseconds (cumulative count 99927)
99.966% <= 1.303 milliseconds (cumulative count 99966)
99.981% <= 1.407 milliseconds (cumulative count 99981)
99.985% <= 1.503 milliseconds (cumulative count 99985)
99.990% <= 1.607 milliseconds (cumulative count 99990)
99.994% <= 1.703 milliseconds (cumulative count 99994)
99.995% <= 1.807 milliseconds (cumulative count 99995)
99.997% <= 2.007 milliseconds (cumulative count 99997)
100.000% <= 2.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 79428.12 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.338     0.136     0.335     0.391     0.559     2.055

## 50 bytes
❯ redis-benchmark -d 50 -t set,get
====== SET ======
  100000 requests completed in 1.30 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.175 milliseconds (cumulative count 1)
50.000% <= 0.335 milliseconds (cumulative count 52673)
75.000% <= 0.359 milliseconds (cumulative count 75995)
87.500% <= 0.399 milliseconds (cumulative count 88273)
93.750% <= 0.479 milliseconds (cumulative count 93775)
96.875% <= 0.687 milliseconds (cumulative count 96920)
98.438% <= 0.983 milliseconds (cumulative count 98457)
99.219% <= 1.359 milliseconds (cumulative count 99222)
99.609% <= 1.631 milliseconds (cumulative count 99632)
99.805% <= 1.655 milliseconds (cumulative count 99823)
99.902% <= 1.687 milliseconds (cumulative count 99918)
99.951% <= 1.847 milliseconds (cumulative count 99952)
99.976% <= 2.159 milliseconds (cumulative count 99976)
99.988% <= 2.263 milliseconds (cumulative count 99988)
99.994% <= 2.319 milliseconds (cumulative count 99994)
99.997% <= 2.407 milliseconds (cumulative count 99997)
99.998% <= 2.559 milliseconds (cumulative count 99999)
99.999% <= 2.615 milliseconds (cumulative count 100000)
100.000% <= 2.615 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.007% <= 0.207 milliseconds (cumulative count 7)
9.160% <= 0.303 milliseconds (cumulative count 9160)
89.256% <= 0.407 milliseconds (cumulative count 89256)
94.534% <= 0.503 milliseconds (cumulative count 94534)
96.237% <= 0.607 milliseconds (cumulative count 96237)
97.040% <= 0.703 milliseconds (cumulative count 97040)
97.577% <= 0.807 milliseconds (cumulative count 97577)
98.053% <= 0.903 milliseconds (cumulative count 98053)
98.529% <= 1.007 milliseconds (cumulative count 98529)
98.822% <= 1.103 milliseconds (cumulative count 98822)
98.987% <= 1.207 milliseconds (cumulative count 98987)
99.141% <= 1.303 milliseconds (cumulative count 99141)
99.287% <= 1.407 milliseconds (cumulative count 99287)
99.405% <= 1.503 milliseconds (cumulative count 99405)
99.528% <= 1.607 milliseconds (cumulative count 99528)
99.935% <= 1.703 milliseconds (cumulative count 99935)
99.947% <= 1.807 milliseconds (cumulative count 99947)
99.956% <= 1.903 milliseconds (cumulative count 99956)
99.962% <= 2.007 milliseconds (cumulative count 99962)
99.970% <= 2.103 milliseconds (cumulative count 99970)
100.000% <= 3.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 76923.08 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.369     0.168     0.335     0.527     1.215     2.615
====== GET ======
  100000 requests completed in 1.41 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.199 milliseconds (cumulative count 5)
50.000% <= 0.351 milliseconds (cumulative count 56623)
75.000% <= 0.391 milliseconds (cumulative count 76211)
87.500% <= 0.495 milliseconds (cumulative count 87565)
93.750% <= 0.599 milliseconds (cumulative count 93795)
96.875% <= 0.767 milliseconds (cumulative count 96941)
98.438% <= 1.079 milliseconds (cumulative count 98459)
99.219% <= 1.519 milliseconds (cumulative count 99236)
99.609% <= 1.727 milliseconds (cumulative count 99614)
99.805% <= 2.071 milliseconds (cumulative count 99806)
99.902% <= 3.839 milliseconds (cumulative count 99903)
99.951% <= 33.343 milliseconds (cumulative count 99952)
99.976% <= 33.919 milliseconds (cumulative count 99976)
99.988% <= 34.303 milliseconds (cumulative count 99988)
99.994% <= 34.495 milliseconds (cumulative count 99994)
99.997% <= 34.623 milliseconds (cumulative count 99997)
99.998% <= 34.687 milliseconds (cumulative count 99999)
99.999% <= 34.719 milliseconds (cumulative count 100000)
100.000% <= 34.719 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.006% <= 0.207 milliseconds (cumulative count 6)
6.360% <= 0.303 milliseconds (cumulative count 6360)
79.179% <= 0.407 milliseconds (cumulative count 79179)
88.180% <= 0.503 milliseconds (cumulative count 88180)
94.089% <= 0.607 milliseconds (cumulative count 94089)
96.250% <= 0.703 milliseconds (cumulative count 96250)
97.222% <= 0.807 milliseconds (cumulative count 97222)
97.801% <= 0.903 milliseconds (cumulative count 97801)
98.217% <= 1.007 milliseconds (cumulative count 98217)
98.529% <= 1.103 milliseconds (cumulative count 98529)
98.751% <= 1.207 milliseconds (cumulative count 98751)
98.897% <= 1.303 milliseconds (cumulative count 98897)
99.015% <= 1.407 milliseconds (cumulative count 99015)
99.199% <= 1.503 milliseconds (cumulative count 99199)
99.419% <= 1.607 milliseconds (cumulative count 99419)
99.578% <= 1.703 milliseconds (cumulative count 99578)
99.708% <= 1.807 milliseconds (cumulative count 99708)
99.759% <= 1.903 milliseconds (cumulative count 99759)
99.790% <= 2.007 milliseconds (cumulative count 99790)
99.811% <= 2.103 milliseconds (cumulative count 99811)
99.900% <= 3.103 milliseconds (cumulative count 99900)
99.915% <= 4.103 milliseconds (cumulative count 99915)
99.949% <= 5.103 milliseconds (cumulative count 99949)
99.950% <= 6.103 milliseconds (cumulative count 99950)
99.981% <= 34.111 milliseconds (cumulative count 99981)
100.000% <= 35.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 70721.36 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.415     0.192     0.351     0.639     1.399    34.719

## 100 bytes                                                                                               
❯ redis-benchmark -d 100 -t set,get
====== SET ======
  100000 requests completed in 1.24 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.175 milliseconds (cumulative count 1)
50.000% <= 0.327 milliseconds (cumulative count 51359)
75.000% <= 0.351 milliseconds (cumulative count 80075)
87.500% <= 0.367 milliseconds (cumulative count 88855)
93.750% <= 0.391 milliseconds (cumulative count 94253)
96.875% <= 0.431 milliseconds (cumulative count 96999)
98.438% <= 0.487 milliseconds (cumulative count 98563)
99.219% <= 0.551 milliseconds (cumulative count 99275)
99.609% <= 0.655 milliseconds (cumulative count 99613)
99.805% <= 0.751 milliseconds (cumulative count 99807)
99.902% <= 0.799 milliseconds (cumulative count 99907)
99.951% <= 0.911 milliseconds (cumulative count 99953)
99.976% <= 1.055 milliseconds (cumulative count 99976)
99.988% <= 1.103 milliseconds (cumulative count 99988)
99.994% <= 1.127 milliseconds (cumulative count 99996)
99.997% <= 1.143 milliseconds (cumulative count 99998)
99.998% <= 1.159 milliseconds (cumulative count 99999)
99.999% <= 1.183 milliseconds (cumulative count 100000)
100.000% <= 1.183 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.004% <= 0.207 milliseconds (cumulative count 4)
13.487% <= 0.303 milliseconds (cumulative count 13487)
95.772% <= 0.407 milliseconds (cumulative count 95772)
98.828% <= 0.503 milliseconds (cumulative count 98828)
99.511% <= 0.607 milliseconds (cumulative count 99511)
99.703% <= 0.703 milliseconds (cumulative count 99703)
99.913% <= 0.807 milliseconds (cumulative count 99913)
99.951% <= 0.903 milliseconds (cumulative count 99951)
99.964% <= 1.007 milliseconds (cumulative count 99964)
99.988% <= 1.103 milliseconds (cumulative count 99988)
100.000% <= 1.207 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 80515.30 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.336     0.168     0.327     0.399     0.519     1.183
====== GET ======
  100000 requests completed in 1.27 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.199 milliseconds (cumulative count 1)
50.000% <= 0.335 milliseconds (cumulative count 54002)
75.000% <= 0.351 milliseconds (cumulative count 75101)
87.500% <= 0.375 milliseconds (cumulative count 90663)
93.750% <= 0.391 milliseconds (cumulative count 94295)
96.875% <= 0.431 milliseconds (cumulative count 97158)
98.438% <= 0.487 milliseconds (cumulative count 98446)
99.219% <= 0.583 milliseconds (cumulative count 99226)
99.609% <= 0.695 milliseconds (cumulative count 99618)
99.805% <= 0.767 milliseconds (cumulative count 99816)
99.902% <= 0.823 milliseconds (cumulative count 99910)
99.951% <= 0.887 milliseconds (cumulative count 99953)
99.976% <= 0.935 milliseconds (cumulative count 99976)
99.988% <= 0.983 milliseconds (cumulative count 99989)
99.994% <= 1.039 milliseconds (cumulative count 99994)
99.997% <= 1.199 milliseconds (cumulative count 99997)
99.998% <= 1.287 milliseconds (cumulative count 99999)
99.999% <= 1.295 milliseconds (cumulative count 100000)
100.000% <= 1.295 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.005% <= 0.207 milliseconds (cumulative count 5)
8.768% <= 0.303 milliseconds (cumulative count 8768)
95.964% <= 0.407 milliseconds (cumulative count 95964)
98.641% <= 0.503 milliseconds (cumulative count 98641)
99.346% <= 0.607 milliseconds (cumulative count 99346)
99.640% <= 0.703 milliseconds (cumulative count 99640)
99.888% <= 0.807 milliseconds (cumulative count 99888)
99.961% <= 0.903 milliseconds (cumulative count 99961)
99.992% <= 1.007 milliseconds (cumulative count 99992)
99.995% <= 1.103 milliseconds (cumulative count 99995)
99.997% <= 1.207 milliseconds (cumulative count 99997)
100.000% <= 1.303 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 78678.20 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.341     0.192     0.335     0.399     0.543     1.295

## 200 bytes
❯ redis-benchmark -d 200 -t set,get
====== SET ======
  100000 requests completed in 1.28 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.175 milliseconds (cumulative count 1)
50.000% <= 0.335 milliseconds (cumulative count 57429)
75.000% <= 0.359 milliseconds (cumulative count 77778)
87.500% <= 0.407 milliseconds (cumulative count 88058)
93.750% <= 0.527 milliseconds (cumulative count 93974)
96.875% <= 0.639 milliseconds (cumulative count 96923)
98.438% <= 0.743 milliseconds (cumulative count 98535)
99.219% <= 0.823 milliseconds (cumulative count 99223)
99.609% <= 0.903 milliseconds (cumulative count 99610)
99.805% <= 0.991 milliseconds (cumulative count 99810)
99.902% <= 1.087 milliseconds (cumulative count 99905)
99.951% <= 1.215 milliseconds (cumulative count 99954)
99.976% <= 1.319 milliseconds (cumulative count 99976)
99.988% <= 1.415 milliseconds (cumulative count 99988)
99.994% <= 1.559 milliseconds (cumulative count 99994)
99.997% <= 1.639 milliseconds (cumulative count 99997)
99.998% <= 1.687 milliseconds (cumulative count 99999)
99.999% <= 1.719 milliseconds (cumulative count 100000)
100.000% <= 1.719 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.012% <= 0.207 milliseconds (cumulative count 12)
11.177% <= 0.303 milliseconds (cumulative count 11177)
88.058% <= 0.407 milliseconds (cumulative count 88058)
93.093% <= 0.503 milliseconds (cumulative count 93093)
96.263% <= 0.607 milliseconds (cumulative count 96263)
98.019% <= 0.703 milliseconds (cumulative count 98019)
99.121% <= 0.807 milliseconds (cumulative count 99121)
99.610% <= 0.903 milliseconds (cumulative count 99610)
99.836% <= 1.007 milliseconds (cumulative count 99836)
99.913% <= 1.103 milliseconds (cumulative count 99913)
99.950% <= 1.207 milliseconds (cumulative count 99950)
99.974% <= 1.303 milliseconds (cumulative count 99974)
99.986% <= 1.407 milliseconds (cumulative count 99986)
99.991% <= 1.503 milliseconds (cumulative count 99991)
99.996% <= 1.607 milliseconds (cumulative count 99996)
99.999% <= 1.703 milliseconds (cumulative count 99999)
100.000% <= 1.807 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 77942.32 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.358     0.168     0.335     0.559     0.791     1.719
====== GET ======
  100000 requests completed in 1.46 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.167 milliseconds (cumulative count 1)
50.000% <= 0.351 milliseconds (cumulative count 52462)
75.000% <= 0.431 milliseconds (cumulative count 75447)
87.500% <= 0.575 milliseconds (cumulative count 87660)
93.750% <= 0.703 milliseconds (cumulative count 93780)
96.875% <= 0.863 milliseconds (cumulative count 96942)
98.438% <= 1.047 milliseconds (cumulative count 98450)
99.219% <= 1.287 milliseconds (cumulative count 99235)
99.609% <= 1.527 milliseconds (cumulative count 99612)
99.805% <= 2.119 milliseconds (cumulative count 99806)
99.902% <= 2.783 milliseconds (cumulative count 99903)
99.951% <= 3.639 milliseconds (cumulative count 99952)
99.976% <= 6.431 milliseconds (cumulative count 99976)
99.988% <= 6.815 milliseconds (cumulative count 99988)
99.994% <= 6.951 milliseconds (cumulative count 99994)
99.997% <= 7.007 milliseconds (cumulative count 99997)
99.998% <= 7.167 milliseconds (cumulative count 99999)
99.999% <= 7.407 milliseconds (cumulative count 100000)
100.000% <= 7.407 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.008% <= 0.207 milliseconds (cumulative count 8)
6.436% <= 0.303 milliseconds (cumulative count 6436)
71.871% <= 0.407 milliseconds (cumulative count 71871)
82.614% <= 0.503 milliseconds (cumulative count 82614)
89.560% <= 0.607 milliseconds (cumulative count 89560)
93.780% <= 0.703 milliseconds (cumulative count 93780)
96.176% <= 0.807 milliseconds (cumulative count 96176)
97.411% <= 0.903 milliseconds (cumulative count 97411)
98.219% <= 1.007 milliseconds (cumulative count 98219)
98.712% <= 1.103 milliseconds (cumulative count 98712)
99.041% <= 1.207 milliseconds (cumulative count 99041)
99.271% <= 1.303 milliseconds (cumulative count 99271)
99.465% <= 1.407 milliseconds (cumulative count 99465)
99.583% <= 1.503 milliseconds (cumulative count 99583)
99.666% <= 1.607 milliseconds (cumulative count 99666)
99.719% <= 1.703 milliseconds (cumulative count 99719)
99.749% <= 1.807 milliseconds (cumulative count 99749)
99.775% <= 1.903 milliseconds (cumulative count 99775)
99.792% <= 2.007 milliseconds (cumulative count 99792)
99.804% <= 2.103 milliseconds (cumulative count 99804)
99.931% <= 3.103 milliseconds (cumulative count 99931)
99.959% <= 4.103 milliseconds (cumulative count 99959)
99.963% <= 5.103 milliseconds (cumulative count 99963)
99.967% <= 6.103 milliseconds (cumulative count 99967)
99.997% <= 7.103 milliseconds (cumulative count 99997)
100.000% <= 8.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 68446.27 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.420     0.160     0.351     0.751     1.199     7.407

## 1000 bytes
❯ redis-benchmark -d 1000 -t set,get
====== SET ======
  100000 requests completed in 1.25 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.175 milliseconds (cumulative count 2)
50.000% <= 0.327 milliseconds (cumulative count 50359)
75.000% <= 0.351 milliseconds (cumulative count 80043)
87.500% <= 0.367 milliseconds (cumulative count 88765)
93.750% <= 0.391 milliseconds (cumulative count 94031)
96.875% <= 0.439 milliseconds (cumulative count 97045)
98.438% <= 0.503 milliseconds (cumulative count 98557)
99.219% <= 0.591 milliseconds (cumulative count 99250)
99.609% <= 0.711 milliseconds (cumulative count 99616)
99.805% <= 0.799 milliseconds (cumulative count 99806)
99.902% <= 0.911 milliseconds (cumulative count 99904)
99.951% <= 1.015 milliseconds (cumulative count 99954)
99.976% <= 1.111 milliseconds (cumulative count 99977)
99.988% <= 1.207 milliseconds (cumulative count 99989)
99.994% <= 1.239 milliseconds (cumulative count 99994)
99.997% <= 1.271 milliseconds (cumulative count 99997)
99.998% <= 1.311 milliseconds (cumulative count 99999)
99.999% <= 1.655 milliseconds (cumulative count 100000)
100.000% <= 1.655 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.007% <= 0.207 milliseconds (cumulative count 7)
12.480% <= 0.303 milliseconds (cumulative count 12480)
95.471% <= 0.407 milliseconds (cumulative count 95471)
98.557% <= 0.503 milliseconds (cumulative count 98557)
99.319% <= 0.607 milliseconds (cumulative count 99319)
99.588% <= 0.703 milliseconds (cumulative count 99588)
99.818% <= 0.807 milliseconds (cumulative count 99818)
99.897% <= 0.903 milliseconds (cumulative count 99897)
99.950% <= 1.007 milliseconds (cumulative count 99950)
99.975% <= 1.103 milliseconds (cumulative count 99975)
99.989% <= 1.207 milliseconds (cumulative count 99989)
99.998% <= 1.303 milliseconds (cumulative count 99998)
99.999% <= 1.407 milliseconds (cumulative count 99999)
100.000% <= 1.703 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 80256.82 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.337     0.168     0.327     0.407     0.551     1.655
====== GET ======
  100000 requests completed in 1.27 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.175 milliseconds (cumulative count 1)
50.000% <= 0.335 milliseconds (cumulative count 54857)
75.000% <= 0.351 milliseconds (cumulative count 75983)
87.500% <= 0.367 milliseconds (cumulative count 87653)
93.750% <= 0.391 milliseconds (cumulative count 94426)
96.875% <= 0.423 milliseconds (cumulative count 97011)
98.438% <= 0.479 milliseconds (cumulative count 98509)
99.219% <= 0.543 milliseconds (cumulative count 99275)
99.609% <= 0.623 milliseconds (cumulative count 99627)
99.805% <= 0.679 milliseconds (cumulative count 99822)
99.902% <= 0.751 milliseconds (cumulative count 99903)
99.951% <= 0.951 milliseconds (cumulative count 99952)
99.976% <= 1.151 milliseconds (cumulative count 99976)
99.988% <= 1.359 milliseconds (cumulative count 99988)
99.994% <= 1.479 milliseconds (cumulative count 99994)
99.997% <= 1.535 milliseconds (cumulative count 99997)
99.998% <= 1.583 milliseconds (cumulative count 99999)
99.999% <= 1.607 milliseconds (cumulative count 100000)
100.000% <= 1.607 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.004% <= 0.207 milliseconds (cumulative count 4)
8.951% <= 0.303 milliseconds (cumulative count 8951)
96.102% <= 0.407 milliseconds (cumulative count 96102)
98.869% <= 0.503 milliseconds (cumulative count 98869)
99.582% <= 0.607 milliseconds (cumulative count 99582)
99.869% <= 0.703 milliseconds (cumulative count 99869)
99.927% <= 0.807 milliseconds (cumulative count 99927)
99.948% <= 0.903 milliseconds (cumulative count 99948)
99.954% <= 1.007 milliseconds (cumulative count 99954)
99.964% <= 1.103 milliseconds (cumulative count 99964)
99.980% <= 1.207 milliseconds (cumulative count 99980)
99.985% <= 1.303 milliseconds (cumulative count 99985)
99.990% <= 1.407 milliseconds (cumulative count 99990)
99.995% <= 1.503 milliseconds (cumulative count 99995)
100.000% <= 1.607 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 78864.35 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.340     0.168     0.335     0.399     0.519     1.607

## 5000 bytes
❯ redis-benchmark -d 5000 -t set,get
====== SET ======
  100000 requests completed in 1.29 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.159 milliseconds (cumulative count 1)
50.000% <= 0.343 milliseconds (cumulative count 58714)
75.000% <= 0.359 milliseconds (cumulative count 76796)
87.500% <= 0.383 milliseconds (cumulative count 89135)
93.750% <= 0.415 milliseconds (cumulative count 94127)
96.875% <= 0.479 milliseconds (cumulative count 97104)
98.438% <= 0.567 milliseconds (cumulative count 98477)
99.219% <= 0.735 milliseconds (cumulative count 99219)
99.609% <= 0.911 milliseconds (cumulative count 99620)
99.805% <= 1.487 milliseconds (cumulative count 99807)
99.902% <= 1.663 milliseconds (cumulative count 99904)
99.951% <= 1.775 milliseconds (cumulative count 99953)
99.976% <= 2.023 milliseconds (cumulative count 99976)
99.988% <= 2.151 milliseconds (cumulative count 99988)
99.994% <= 2.527 milliseconds (cumulative count 99994)
99.997% <= 2.583 milliseconds (cumulative count 99997)
99.998% <= 2.615 milliseconds (cumulative count 99999)
99.999% <= 2.647 milliseconds (cumulative count 100000)
100.000% <= 2.647 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.007% <= 0.207 milliseconds (cumulative count 7)
4.379% <= 0.303 milliseconds (cumulative count 4379)
93.378% <= 0.407 milliseconds (cumulative count 93378)
97.662% <= 0.503 milliseconds (cumulative count 97662)
98.709% <= 0.607 milliseconds (cumulative count 98709)
99.132% <= 0.703 milliseconds (cumulative count 99132)
99.413% <= 0.807 milliseconds (cumulative count 99413)
99.605% <= 0.903 milliseconds (cumulative count 99605)
99.705% <= 1.007 milliseconds (cumulative count 99705)
99.734% <= 1.103 milliseconds (cumulative count 99734)
99.750% <= 1.207 milliseconds (cumulative count 99750)
99.758% <= 1.303 milliseconds (cumulative count 99758)
99.763% <= 1.407 milliseconds (cumulative count 99763)
99.811% <= 1.503 milliseconds (cumulative count 99811)
99.874% <= 1.607 milliseconds (cumulative count 99874)
99.933% <= 1.703 milliseconds (cumulative count 99933)
99.958% <= 1.807 milliseconds (cumulative count 99958)
99.965% <= 1.903 milliseconds (cumulative count 99965)
99.975% <= 2.007 milliseconds (cumulative count 99975)
99.984% <= 2.103 milliseconds (cumulative count 99984)
100.000% <= 3.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 77220.08 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.353     0.152     0.343     0.431     0.671     2.647
====== GET ======
  100000 requests completed in 1.31 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.199 milliseconds (cumulative count 2)
50.000% <= 0.343 milliseconds (cumulative count 53124)
75.000% <= 0.367 milliseconds (cumulative count 81224)
87.500% <= 0.383 milliseconds (cumulative count 89598)
93.750% <= 0.407 milliseconds (cumulative count 94615)
96.875% <= 0.447 milliseconds (cumulative count 97096)
98.438% <= 0.503 milliseconds (cumulative count 98536)
99.219% <= 0.567 milliseconds (cumulative count 99253)
99.609% <= 0.663 milliseconds (cumulative count 99621)
99.805% <= 0.783 milliseconds (cumulative count 99805)
99.902% <= 0.927 milliseconds (cumulative count 99905)
99.951% <= 1.095 milliseconds (cumulative count 99952)
99.976% <= 1.247 milliseconds (cumulative count 99976)
99.988% <= 1.343 milliseconds (cumulative count 99988)
99.994% <= 1.471 milliseconds (cumulative count 99994)
99.997% <= 1.551 milliseconds (cumulative count 99997)
99.998% <= 1.599 milliseconds (cumulative count 99999)
99.999% <= 1.631 milliseconds (cumulative count 100000)
100.000% <= 1.631 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.003% <= 0.207 milliseconds (cumulative count 3)
4.281% <= 0.303 milliseconds (cumulative count 4281)
94.615% <= 0.407 milliseconds (cumulative count 94615)
98.536% <= 0.503 milliseconds (cumulative count 98536)
99.446% <= 0.607 milliseconds (cumulative count 99446)
99.701% <= 0.703 milliseconds (cumulative count 99701)
99.829% <= 0.807 milliseconds (cumulative count 99829)
99.893% <= 0.903 milliseconds (cumulative count 99893)
99.934% <= 1.007 milliseconds (cumulative count 99934)
99.953% <= 1.103 milliseconds (cumulative count 99953)
99.968% <= 1.207 milliseconds (cumulative count 99968)
99.984% <= 1.303 milliseconds (cumulative count 99984)
99.992% <= 1.407 milliseconds (cumulative count 99992)
99.994% <= 1.503 milliseconds (cumulative count 99994)
99.999% <= 1.607 milliseconds (cumulative count 99999)
100.000% <= 1.703 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 76569.68 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.350     0.192     0.343     0.415     0.543     1.631
