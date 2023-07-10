import json

def decode_chinese(json_data):
    # 读取JSON数据
    with open(json_data, 'r', encoding='utf-8') as file:
        data = json.load(file)

    # 将Unicode编码转换为普通中文字符串
    decoded_data = json.dumps(data, ensure_ascii=False)

    # 写入输出文件
    with open('output.json', 'w', encoding='utf-8') as file:
        file.write(decoded_data)

# 使用示例
json_file = 'example1.json'
decode_chinese(json_file)
