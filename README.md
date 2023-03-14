# SnapFoodTask Detail

A ) Main challenge :  
 we want to write an app to manage order process  
design a restful api that expose 1 endpoint  
POST api/order input : {“order_id”:10,”price”:1000,”title”:”burger”}  
in the receiving point send the received data to redis (queue)  
create another app to read data from the above redis and process the order  
the process should goes like this :  
use mysql to save the retrieved data into orders table  

    Acceptance criteria :
        1 – use multi layer config (file,environment,…)
        2 – use docker compose for external services (mysql,redis)
        2 – include build.sh file to build the apps
        3 – include run.sh to run the apps
        4 – writing unit tests is a plus

B ) suppose we have slice of integers which every element is >=1 and all the numbers  
are repeated twice except for 1 number for example : {2,2,5,6,5} we want a way to  
find the exception number with only 1 iteration (loop) in this case the answer is 6  

C ) write a function that takes slice of urls as input and request each url concurrently
to get the response  
use 1 second timeout foreach request after aggregate all responses  
print the responses to console if all the calls were successful  

D ) we have var called nums = []int{12,54,89,21,66,47,14,285,96,…}  
we want to calculate the sum of each number by the power of 2 (12^2+54^2+..)  
we want a pipeline to send this numbers to a channel and have 2 concurrent functions  
(worker) receives from the pipeline and do the calc  

---

# How Run Tasks

1 - For task A read readme file in A Question

2 - For Run task B/C/D"
   - go to  " B/C/D Question " folder,
   run main.go < ` go run main ` >, 
   this is very simple test code, 
   for more test run < ` go test ./... ` >.



### Worng
This Repo After Interview Changed !  
My First Answer Was 2 month before last change!