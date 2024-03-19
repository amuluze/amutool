# -*- coding: utf-8 -*-
"""
@ Author     : Amu
@ Date       : 2023/11/30 16:57:56
@ Description:
"""

from kafka3 import KafkaClient


client = KafkaClient(bootstrap_servers=["localhost:9092"])
print(client.check_version())
