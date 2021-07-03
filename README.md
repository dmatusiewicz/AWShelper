# AWShelper 

Helper for getting mfa session or role. 


Installation (assuming you have Go and Git installed): 
```
git clone git@github.com:dmatusiewicz/AWShelper.git
cd ./AWShelper/cmd/AWShelper; go install 
```

Example usage: 
```
eval $(AWShelper session -m <MFA_TOKEN>)
eval $(AWShelper role -r admin -m <MFA_TOKEN>)
```


