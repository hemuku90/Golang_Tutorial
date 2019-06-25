class node:
    def __init__(self,data=None):
        self.data=data
        self.next=None
class linkedlist:
    def __init__(self):
        self.head=node() 

    def append(self,data):
        new_node=node(data)  
        self.data=data
        cur=self.head 
        while cur.next!=None:
            cur=cur.next
        cur.next=new_node
    
    def length(self):
        cur=self.head
        total=0
        while cur.next!=None:
            total+=1
            cur=cur.next
        return total
    def display(self):
        elems=[]
        cur_node=self.head
        while cur_node.next!=None:
            cur_node=cur_node.next
            elems.append(cur_node.data)
        return elems
    def get(self,index):
        if index>=self.length():
            print("Error: index out of bound")
            return None
        cur_index=0
        cur_node=self.head
        while True:
            cur_node=cur_node.next
            if cur_index==index:return cur_node.data
            cur_index+=1

mylist=linkedlist()
for i in range(1,6,1):
    mylist.append(i)
print(mylist.display())
print(mylist.get(2))

