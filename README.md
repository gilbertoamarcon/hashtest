# HashTest

A golang hashing method test.

This test creates a set of randomly generated structs and tests multiple hashing methods that turn the structs into a fixed-size string hash.

The test evaluates the total time duration taken by each method, whether the method succeeded in generating collision-free hashes, and the total size of the hash set. 

Example usage:
```
./build_and_run.sh
``` 

Results:
```
StructhashMD5: true 2.466313764s 16000000
Fnv32: false 832.039213ms 4000000
Fnv32a: false 848.506689ms 4000000
Fnv64: true 853.616123ms 8000000
Fnv64a: true 902.344245ms 8000000
Fnv128: true 952.440633ms 16000000
Fnv128a: true 926.625991ms 16000000
Sha1: true 1.062828202s 20000000
Sha3: true 1.649616463s 32000000
Sha256: true 1.165236608s 32000000
Md4: true 1.280617254s 16000000
Md5: true 963.332718ms 16000000
Blake2b: true 1.000788289s 32000000
```
