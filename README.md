# Afanty [![Build Status](https://travis-ci.org/zhyuri/afanty.svg?branch=master)](https://travis-ci.org/zhyuri/afanty)

Afanty is a microservice orchestration service based on Go.

In my daily develop work, there are a lot of `business logic` code in the backend.
Sometimes it's just read Database and transform the data into some data structure and response. Sometimes it makes some
RPC call and combine the result into the response. After coding for a while, I realized that there must be a solution
to avoid writing such 'idiot' code. Even maybe one day the AI can do such repeated work, because the 'business logic' is
another kind of pattern for feature description.

Nowadays people are talking about microservice architecture and infrastructure support it. With a RPC framework you can
combine many sub-systems together to accomplish a greater feature. But there are still some code to write time after time,
just for different feature. And you will find out most of your code is duplicate.
