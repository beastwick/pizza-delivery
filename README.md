# Pizza Delivery

See Problem.md for the scenario.

## About

While I am a PHP developer, I take seriously working in other languages and have done so throughout my career. I've decided to do this project in Go to demonstrate that I can learn quickly.

## Run

Run the main binary in main/ and pass the input file.

```
$ cd main/
$ ./main ../PizzaDeliveryInput.txt
```

## Problem Case Output

Please see PizzaDeliveryOutput.txt for the full log.

>Route: 1  
>1 deliverer(s) visited 2565 homes and delivered 8193 pizzas.  

>Route: 1    
>2 deliverer(s) visited 2639 homes and delivered 8194 pizzas.  

## Test Case Output

>Route: 1    
>Directions: [>]  
>1 deliverer(s) visited 2 homes and delivered 2 pizzas.  

>Route: 2  
>Directions: [^ v]  
>1 deliverer(s) visited 2 homes and delivered 3 pizzas.  

>Route: 3  
>Directions: [^ > v <]  
>1 deliverer(s) visited 4 homes and delivered 5 pizzas.  

>Route: 4  
>Directions: [^ v ^ v ^ v ^ v ^ v]  
>1 deliverer(s) visited 2 homes and delivered 11 pizzas.  

>Route: 1  
>Directions: [>]  
>2 deliverer(s) visited 2 homes and delivered 3 pizzas.  

>Route: 2  
>Directions: [^ v]  
>2 deliverer(s) visited 3 homes and delivered 4 pizzas.  

>Route: 3  
>Directions: [^ > v <]  
>2 deliverer(s) visited 3 homes and delivered 6 pizzas.  

>Route: 4  
>Directions: [^ v ^ v ^ v ^ v ^ v]  
>2 deliverer(s) visited 11 homes and delivered 12 pizzas.  

## Extra

I generated a small call graph found at main/cgraph.pdf and although not very useful here, I find they help me orient myself in new projects and are fun to look at. Just note this graph goes too deep because of library dependencies, so just glancing at the top helps. 
