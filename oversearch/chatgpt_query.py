import sys
import openai
import json

def main(query):
    openai.api_key = 'your_openai_api_key'

    response = openai.Completion.create(
        engine="text-davinci-003",
        prompt=query,
        max_tokens=100
    )

    result = {
        "query": query,
        "response": response.choices[0].text.strip()
    }

    print(json.dumps(result))

if __name__ == "__main__":
    query = sys.argv[1]
    main(query)
