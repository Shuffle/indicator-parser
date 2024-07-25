import ctypes
import os
import json
import ipaddress
from enum import Enum

lib_path = os.path.abspath(os.path.join(
    os.path.dirname(__file__), "..", "build", "lib.so"))
ioc_lib = ctypes.CDLL(lib_path)


class IndicatorType(str, Enum):
    IPV4 = "IPV4"
    URL_LINK = "URL_LINK"
    Domain = "DOMAIN"
    Email = "EMAIL"


ioc_lib.ParseIOC.argtypes = [ctypes.c_char_p, ctypes.c_char_p]
ioc_lib.ParseIOC.restype = ctypes.c_char_p


def parse(value, types):
    types_json = json.dumps([t.value for t in types])
    result_ptr = ioc_lib.ParseIOC(value.encode(
        'utf-8'), types_json.encode('utf-8'))
    json_str = ctypes.string_at(result_ptr).decode('utf-8')

    ctypes.CDLL("libc.so.6").free(result_ptr)

    result = json.loads(json_str)
    return result


def is_private_ip(ip_string):
    try:
        ip_obj = ipaddress.ip_address(ip_string)
        return ip_obj.is_private
    except ValueError:
        return False


def parse_ioc(self, input_string, input_type="all"):
    if input_type == "" or input_type == "all":
        types_to_use = list(IndicatorType)
    else:
        input_type = input_type.split(",")
        input_type = [item.strip().upper() for item in input_type]
        types_to_use = [IndicatorType(
            t) for t in input_type if t in IndicatorType.__members__]

    parsed_results = parse(input_string, types_to_use)

    newarray = []
    for item in parsed_results:
        data_type = item['type'].lower()
        if data_type == "url_link":
            data_type = "url"

        data = {
            "data": item['data'],
            "data_type": data_type
        }

        if data_type == "ipv4":
            data["data_type"] = "ip"
            data["is_private_ip"] = is_private_ip(item['data'])

        if data not in newarray:
            newarray.append(data)

    try:
        return json.dumps(newarray)
    except json.JSONDecodeError as e:
        return f"Failed to parse IOC's: {e}"
