function Node(element) {
  this.element = element;
  this.next = null;
}

function LinkedList() {
  this.head = new Node("head");
  this.find = find;
  this.insert = insert;
  this.remove = remove;
  this.display = display;
  this.findPrevious = findPrevious;
}

// find element in list
function find(item) {
  let currentNode = this.head;
  while (currentNode.element != item) {
    currentNode = currentNode.next;
  }
  return currentNode;
}

// insert new node after posItem
function insert(newElement, posItem) {
  let newNode = new Node(newElement);
  let currentNode = this.find(posItem);
  newNode.next = currentNode.next;
  currentNode.next = newNode;
}

function display() {
  let str = "";
  let currentNode = this.head;
  while (currentNode.next != null) {
    if (currentNode.next.next != null) {
      str += `${currentNode.next.element} -> `;
    } else {
      str += `${currentNode.next.element}`;
    }
    currentNode = currentNode.next;
  }
  console.log(str);
}

function findPrevious(item) {
  let currentNode = this.head;
  while (currentNode.next != null && currentNode.next.element != item) {
    currentNode = currentNode.next;
  }
  return currentNode;
}

function remove(item) {
  let previousNode = this.findPrevious(item);
  if (previousNode.next != null) {
    previousNode.next = previousNode.next.next;
  }
}

var cities = new LinkedList();
cities.insert("Beijing", "head");
cities.insert("Shanghai", "Beijing");
cities.insert("Shenzhen", "Shanghai");
cities.insert("Guangzhou", "Shenzhen");
cities.display();
cities.remove("Shenzhen");
cities.display();
cities.remove("Beijing");
cities.display();
cities.insert("Hangzhou", "Shanghai");
cities.display();
