import boto3

region = 'us-east-1'

instances = ['i-0a8e80db6e9db71d6','i-0b10a7ce9c0e46393']
#instance = 'i-0a8e80db6e9db71d6'

def lambda_handler(event, context):
    ec2 = boto3.client('ec2', region_name=region)
    ec2.start_instances(InstanceIds=instances)
    #startingUp = ec2.instances.filter(InstanceIds=instance).start()

print "started your instances: " + str(instances)