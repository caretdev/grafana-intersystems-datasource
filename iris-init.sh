#!/bin/bash

iris session iris -U%SYS '##class(Security.Users).UnExpireUserPasswords("*")'