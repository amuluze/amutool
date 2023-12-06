#!/usr/bin/env python
# -*- coding: utf-8 -*-
"""
@Time    : 2023/11/30 14:41
@Author  : Amu
@Desc    :
"""
import json
import random

from kafka3 import KafkaProducer

if __name__ == '__main__':
    producer = KafkaProducer(bootstrap_servers=["localhost:9092"])
    topic = "test"
    for i in range(100000):
        msg = {
            "username": "jack-" + str(random.randint(0, 1000000)),
            "age": random.randint(0, 120),
            "sex": 1
        }
        msg_str = json.dumps(msg)
        print(msg_str)
        res = producer.send(topic, msg_str.encode("utf-8"))
        print(res.value)
