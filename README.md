# IAM Role JSON Verifier

This is a simple tool written in Go to verify AWS IAM Role JSON data. It checks if the "Resource" field in the input JSON contains a single asterisk (*) and returns `false` in that case, otherwise it returns `true`.

## How to Use

1. Clone the repository:
```sh
git clone <repository_url>
```

2. Navigate to the project directory:
```sh
cd IAM-Role-JSON-Verifier
```

3. Build the application:
```sh
go build
```

4. Run the program with your JSON file:

- For Linux/macOS:: 
```sh
./policyvalidator <path_to_your_json_file>
```

- For Windows: 
```sh
./policyvalidator.exe <path_to_your_json_file>
```


Replace <path_to_your_json_file> with the path to your JSON file containing the IAM Role policy.

5. The program will output `true` if the JSON data is valid, and `false` otherwise.

## Example

Suppose you have a JSON file named `policy.json` containing the following IAM Role JSON data:
```json
{
    "PolicyName": "root",
    "PolicyDocument": {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "IamListAccess",
                "Effect": "Allow",
                "Action": [
                    "iam:ListRoles",
                    "iam:ListUsers"
                ],
                "Resource": "lorem:ipsum"
            }
        ]
    }
}
```

You can verify this JSON data by running:

- For Linux/macOS:: 
```sh
./policyvalidator policy.json
```

- For Windows: 
```sh
./policyvalidator.exe policy.json
```

## Tests

To run the tests, execute the following command:
```sh
go test
```

## JSON Data Format

The expected format of the input JSON data should be as follows:

```json
{
  "PolicyDocument": {
    "Statement": [
      {
        "Resource": "string or [string]"
      }
    ]
  }
}
```

- `Resource`: Specifies the Amazon Resource Name (ARN) of the resource or resources to which the action applies.
