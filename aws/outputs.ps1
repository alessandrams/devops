 aws iam create-role --role-name user-lambda-ec2 --assume-role-policy-document file://policy.json

 ############################################ OUTPUT ###########################################
{
    "Role": {
        "Path": "/",
        "RoleName": "user-lambda",
        "RoleId": "AROA5TVTMOHQ4RXK5247D",
        "Arn": "arn:aws:iam::935605334497:role/user-lambda",
        "CreateDate": "2019-08-07T20:48:14Z",
        "AssumeRolePolicyDocument": {
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Effect": "Allow",
                    "Principal": {
                        "Service": "lambda.amazonaws.com"
                    },
                    "Action": "sts:AssumeRole"
                }
            ]
        }
    }
}
############################################ OUTPUT ###########################################

aws lambda create-function --function-name start-ec2 --zip-file fileb://lambda.zip --handler lambda.lambda_handler --runtime python2.7 --role arn:aws:iam::935605334497:role/user-lambda

{
    "FunctionName": "start-ec2",
    "FunctionArn": "arn:aws:lambda:us-east-1:935605334497:function:start-ec2",
    "Runtime": "python2.7",
    "Role": "arn:aws:iam::935605334497:role/user-lambda",
    "Handler": "lambda.lambda_handler",
    "CodeSize": 358,
    "Description": "",
    "Timeout": 3,
    "MemorySize": 128,
    "LastModified": "2019-08-07T20:54:27.175+0000",
    "CodeSha256": "VDMwuxjSywDGudcgQSPFLFZXci1IUQmbH1mHPoBeYT4=",
    "Version": "$LATEST",
    "TracingConfig": {
        "Mode": "PassThrough"
    },
    "RevisionId": "8d2949c3-2adc-4b75-ba33-5ea4afbf6b8e"
}

aws lambda update-function-configuration --function-name start-ec2 --timeout 45

############################################ OUTPUT ###########################################

aws lambda invoke --function-name start-ec2 myresults.txt

{
    "StatusCode": 200,
    "FunctionError": "Unhandled",
    "ExecutedVersion": "$LATEST"
}

aws lambda delete-function --function-name start-ec2

############################################ OUTPUT ###########################################





############################################ OUTPUT ###########################################




     ############################################ OUTPUT ###########################################