def divide(weight):
    if(weight<=2 or (weight%2 != 0)):
        print("NO")
    else:
        print("YES")

weight = int(input())
divide(weight)