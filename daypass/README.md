## Description

DAYPASS is a package used to secure data with dynamic encryption. For example, it is used for passwords that will automatically change the hashing algorithm based on the creation date.

For instance, a password created on July 1st will use the Argon2 hashing algorithm, while a password created on August 2nd will use the SHA-512 hashing algorithm.