### For problem 1 Two-Sum

Add later 

### For problem 53

#### Example Walkthrough (nums = [-2,1,-3,4,-1,2,1,-5,4]):
1. Initialize `maxSum = -2`, `currentSum = -2` (first element).
- Iteration 1 (i = 1): `nums[1] = 1`
  - `currentSum = max(1, -2 + 1) = max(1, -1) = 1`
  - `maxSum = max(-2, 1) = 1`
- Iteration 2 (i = 2): `nums[2] = -3`
  - `currentSum = max(-3, 1 + (-3)) = max(-3, -2) = -2`
  - `maxSum = max(1, -2) = 1`
- Iteration 3 (i = 3): `nums[3] = 4`
  - `currentSum = max(4, -2 + 4) = max(4, 2) = 4`
  - `maxSum = max(1, 4) = 4`
- Iteration 4 (i = 4): nums[4] = -1
  - currentSum = max(-1, 4 + (-1)) = max(-1, 3) = 3
  - maxSum = max(4, 3) = 4
- Iteration 5 (i = 5): nums[5] = 2
  - currentSum = max(2, 3 + 2) = max(2, 5) = 5
  - maxSum = max(4, 5) = 5
- Iteration 6 (i = 6): nums[6] = 1
  - currentSum = max(1, 5 + 1) = max(1, 6) = 6
  - maxSum = max(5, 6) = 6
- Iteration 7 (i = 7): nums[7] = -5
  - currentSum = max(-5, 6 + (-5)) = max(-5, 1) = 1
  - maxSum = max(6, 1) = 6
- Iteration 8 (i = 8): nums[8] = 4
  - currentSum = max(4, 1 + 4) = max(4, 5) = 5
  - maxSum = max(6, 5) = 6
- After all iterations, the `maxSum` is `6`, which is the largest sum for the - subarray `[4, -1, 2, 1]`.

### For problem 3 Longest Substring Without Repeating Characters

Let's walk through the algorithm with the string "abcbabcd" to determine the length of the longest substring without repeating characters.

#### **Initialization:**
- charIndexMap = {} (empty)
- start = 0
- maxLength = 0

#### **Processing the string "abcbabcd":**
1. **Iteration 1:** end = 0, char = 'a'
- charIndexMap does not have 'a', so no update to start.
- Update charIndexMap to { 'a': 0 }
- currentLength = end - start + 1 = 0 - 0 + 1 = 1
- maxLength is updated to 1


2. **Iteration 2:** end = 1, char = 'b'
- charIndexMap does not have 'b', so no update to start.
- Update charIndexMap to { 'a': 0, 'b': 1 }
- currentLength = end - start + 1 = 1 - 0 + 1 = 2
- maxLength is updated to 2


3. **Iteration 3:** end = 2, char = 'c'
- charIndexMap does not have 'c', so no update to start.
- Update charIndexMap to { 'a': 0, 'b': 1, 'c': 2 }
- currentLength = end - start + 1 = 2 - 0 + 1 = 3
- maxLength is updated to 3


4. **Iteration 4:** end = 3, char = 'b'
- charIndexMap has 'b' at index 1, which is >= start.
- Update start to index + 1 = 1 + 1 = 2
- Update charIndexMap to { 'a': 0, 'b': 3, 'c': 2 }
- currentLength = end - start + 1 = 3 - 2 + 1 = 2
- maxLength remains 3


5. **Iteration 5:** end = 4, char = 'a'
- charIndexMap has 'a' at index 0, which is < start.
- No update to start.
- Update charIndexMap to { 'a': 4, 'b': 3, 'c': 2 }
- currentLength = end - start + 1 = 4 - 2 + 1 = 3
- maxLength remains 3


6. **Iteration 6:** end = 5, char = 'b'
charIndexMap has 'b' at index 3, which is >= start.
- Update start to index + 1 = 3 + 1 = 4
- Update charIndexMap to { 'a': 4, 'b': 5, 'c': 2 }
- currentLength = end - start + 1 = 5 - 4 + 1 = 2
- maxLength remains 3


7. **Iteration 7:** end = 6, char = 'c'
- charIndexMap has 'c' at index 2, which is < start.
- No update to start.
- Update charIndexMap to { 'a': 4, 'b': 5, 'c': 6 }
- currentLength = end - start + 1 = 6 - 4 + 1 = 3
- maxLength remains 3


8. **Iteration 8:** end = 7, char = 'd'
- charIndexMap does not have 'd', so no update to start.
- Update charIndexMap to { 'a': 4, 'b': 5, 'c': 6, 'd': 7 }
- currentLength = end - start + 1 = 7 - 4 + 1 = 4
- maxLength is updated to 4

### For problem 198 House Robber

Let’s walk through the example [1, 3, 2, 1]:

#### Initialization:
- prev1 = 0
- prev2 = 0
#### Iterate Through Houses:
- For the first house (1):
- - newPrev = max(prev1 + 1, prev2) = max(0 + 1, 0) = 1
- Update: prev1 = 0, prev2 = 1
- - For the second house (3):
- - newPrev = max(prev1 + 3, prev2) = max(0 + 3, 1) = 3
- Update: prev1 = 1, prev2 = 3
- - For the third house (2):
- - newPrev = max(prev1 + 2, prev2) = max(1 + 2, 3) = 3
- Update: prev1 = 3, prev2 = 3
- - For the fourth house (1):
- - newPrev = max(prev1 + 1, prev2) = max(3 + 1, 3) = 4
- Update: prev1 = 3, prev2 = 4
#### Result:
- The result is prev2 = 4

### For problem 213 House Robber 2 

Let’s walk through the example [2, 3, 2]:

#### Exclude the first house (2):
- We consider the array [3, 2].
- Applying the robLinear function:
- - Initialize prev1 = 0 and prev2 = 0.
- - For 3: newPrev = max(0 + 3, 0) = 3. Update: prev1 = 0, prev2 = 3.
- - For 2: newPrev = max(0 + 2, 3) = 3. Update: prev1 = 3, prev2 = 3.
- The maximum amount for this case is 3.
#### Exclude the last house (2):
- We consider the array [2, 3].
- Applying the robLinear function:
- - Initialize prev1 = 0 and prev2 = 0.
- - For 2: newPrev = max(0 + 2, 0) = 2. Update: prev1 = 0, prev2 = 2.
- - For 3: newPrev = max(0 + 3, 2) = 3. Update: prev1 = 2, prev2 = 3.
- The maximum amount for this case is 3.
#### Result:
- The maximum amount considering both cases is max(3, 3) = 3.

### Leetcode 33 Search in Rotated Sorted Array

#### Explanation with Example [4, 5, 6, 7, 0, 1, 2] and Target 0
- **Initialization:**
- - left = 0
- - right = 6 (length of array - 1)

- **First Iteration:**
- - mid = (0 + 6) / 2 = 3
- - nums[mid] = 7
- - nums[left] <= nums[mid] (4 <= 7): Left part [4, 5, 6, 7] is sorted.
- - target = 0 is not in the range [4, 7]
- - Search the right part, so update left = mid + 1 = 4

- **Second Iteration:**
- - mid = (4 + 6) / 2 = 5
- - nums[mid] = 1
- - nums[left] <= nums[mid] (0 <= 1): Left part [0, 1] is sorted.
- - target = 0 is in the range [0, 1]
- - Search the left part, so update right = mid - 1 = 4

- **Third Iteration:**
- - mid = (4 + 4) / 2 = 4
- - nums[mid] = 0
- - nums[mid] == target, so return mid = 4

### Find Minimum in Rotated Sorted Array

#### Explanation with Example [3, 4, 5, 1, 2]
**Initialization:**
- left = 0
- right = 4 (last index of the array)

**First Iteration:**
- mid = (0 + 4) / 2 = 2
- nums[mid] = 5
- Compare nums[mid] (5) with nums[right] (2): 5 > 2
- - Minimum value is in the right part of the array (excluding - mid).
- - Update left = mid + 1 = 3

**Second Iteration:**
- mid = (3 + 4) / 2 = 3
- nums[mid] = 1
- Compare nums[mid] (1) with nums[right] (2): 1 < 2
- - Minimum value is in the left part of the array (including mid).
- - Update right = mid = 3

**Termination:**
- The loop ends when left equals right (left = right = 3).
- The minimum value is nums[left], which is 1.


### Leetcode 15 3 Sum

#### Explanation with Example [-1, 0, 1, 2, -1, -4]
1. Sorting the Array:
- Sorted array: [-4, -1, -1, 0, 1, 2]
2. Iterate and Fix One Element:
- Fix -4 (i=0):
- - The target sum for the two-pointer search is 4 (0 - (-4)).
- - Left pointer starts at index 1 (-1), right pointer starts at index 5 (2).
- - Sum = -4 + (-1) + 2 = -3 (Less than 0, move left pointer to index 2)
- - Sum = -4 + (-1) + 2 = -3 (Still less, move left pointer to index 3)
- - Sum = -4 + 0 + 2 = -2 (Still less, move left pointer to index 4)
- - Sum = -4 + 1 + 2 = -1 (Still less, move left pointer to index 5)
- - No valid triplet found for -4, move to the next element.
- Fix -1 (i=1):
- - The target sum for the two-pointer search is 1 (0 - (-1)).
- - Left pointer starts at index 2 (-1), right pointer starts at index 5 (2).
- - Sum = -1 + (-1) + 2 = 0 (Found a triplet [-1, -1, 2])
- - Skip duplicates for -1 and 2, move both pointers.
- - Next pointers: Left at index 4 (1), right at index 4 (same, end loop).
- - Triplet [-1, -1, 2] added.
- Fix -1 (i=2):
- - Skip as -1 at index 2 is a duplicate of -1 at index 1.
- Fix 0 (i=3):
- - The target sum for the two-pointer search is 0 - 0 = 0.
- - Left pointer starts at index 4 (1), right pointer starts at index 5 (2).
- - Sum = 0 + 1 + 2 = 3 (Greater than 0, move right pointer to index 4)
- - No valid triplet found for 0.
3. Final Result:
- The unique triplets that sum up to zero are [-1, -1, 2] and [-1, 0, 1].

### Leetcode 56 Merge Intervals

#### Explanation with Example [[1, 3], [2, 6], [8, 10], [15, 18]]
1. Sort the Intervals:
- The intervals are already sorted by their start values: [[1, 3], [2, 6], [8, 10], [15, 18]].
2. Merge Intervals:
- Start with the first interval [1, 3] and initialize it as currentInterval.
- Process [2, 6]:
- - Overlaps with [1, 3] because 2 <= 3. Merge them by updating currentInterval to [1, 6] (the maximum end value between 3 and 6).
- Process [8, 10]:
- - Does not overlap with [1, 6] because 8 > 6. Add [1, 6] to the result and update currentInterval to [8, 10].
- Process [15, 18]:
- - Does not overlap with [8, 10] because 15 > 10. Add [8, 10] to the result and update currentInterval to [15, 18].
3. Add the Last Interval:
- Add the last currentInterval [15, 18] to the result.
4. Result:
- The merged intervals are [[1, 6], [8, 10], [15, 18]].

### Leetcode 57 - Insert Interval
**Input:**
- Intervals: [[1, 3], [6, 9]]
- New Interval: [2, 5]
**Steps:**

1. Add New Interval:
- Intervals after adding new interval: [[1, 3], [6, 9], [2, 5]]
2. Sort the Intervals:
- After sorting by start times: [[1, 3], [2, 5], [6, 9]]
3. Merge Intervals:
- Start with [1, 3] as currentInterval.
- Compare with [2, 5]: They overlap because 2 <= 3. Merge to form [1, 5].
- Compare with [6, 9]: No overlap with [1, 5]. Add [1, 5] to the result and update currentInterval to [6, 9].
- Add the last currentInterval [6, 9] to the result.
**Result:**
- Merged Intervals: [[1, 5], [6, 9]]

### Leetcode 435 - Non-overlapping Intervals

**Input:**
- Intervals: [[1, 2], [2, 3], [3, 4], [1, 3]]

**Steps:**
1. Sort Intervals by End Time:
- The intervals sorted by end times are [[1, 2], [2, 3], [1, 3], [3, 4]].
2. Iterate and Select Non-Overlapping Intervals:
- Initialize count as 1 (select the first interval [1, 2]).
- Set end to 2 (end time of the first interval).
- Compare the next interval [2, 3]:
- - Start of [2, 3] is equal to end, so it overlaps but can be included. Update end to 3 and increment count.
- Compare the next interval [1, 3]:
- - Start of [1, 3] is less than end (3), so it overlaps. Skip this interval.
- Compare the last interval [3, 4]:
- - Start of [3, 4] is equal to end, so it can be included. Update end to 4 and increment count.
3. Calculate the Number of Intervals to Remove:
- Total number of intervals is 4.
- Number of non-overlapping intervals is 3 ([[1, 2], [2, 3], [3, 4]]).
- Number of intervals to remove is 4 - 3 = 1.

**Result:**
- The minimum number of intervals to remove is 1.

### Leetcode 206 - 

**Input Linked List:**

- 1 -> 2 -> 3 -> 4 -> nil
**Steps:**

1. **Initialization:**
- prev = nil
- current = 1 -> 2 -> 3 -> 4 -> nil
2. **Iteration and Reversal:**
- **First Iteration:**
- - next = 2 -> 3 -> 4 -> nil
- - current.Next = nil (reversed pointer)
- - Move prev to 1 -> nil
- - Move current to 2 -> 3 -> 4 -> nil
- **Second Iteration:**
- - next = 3 -> 4 -> nil
- - current.Next = 1 -> nil
- - Move prev to 2 -> 1 -> nil
- - Move current to 3 -> 4 -> nil
- **Third Iteration:**
- - next = 4 -> nil
- - current.Next = 2 -> 1 -> nil
- - Move prev to 3 -> 2 -> 1 -> nil
- - Move current to 4 -> nil
- **Fourth Iteration:**
- - next = nil
- - current.Next = 3 -> 2 -> 1 -> nil
- - Move prev to 4 -> 3 -> 2 -> 1 -> nil
- - Move current to nil
3. **Completion:**
- The prev pointer now points to the new head of the reversed list: 4 -> 3 -> 2 -> 1 -> nil.

```
func printList(head *ListNode) {
    for head != nil {
        fmt.Print(head.Val, " -> ")
        head = head.Next
    }
    fmt.Println("nil")
}
```

1. **Initial Call:**
go
Copy code
```
printList(head) // head points to the node with value 1
```
2. **First Iteration:**
- head points to the node with value 1.
- Print 1 ->.
- Move head to the node with value 2.
3. **Second Iteration:**
- head points to the node with value 2.
- Print 2 ->.
- Move head to the node with value 3.
4. **Third Iteration:**
head points to the node with value 3.
- Print 3 ->.
- Move head to the node with value 4.
5. **Fourth Iteration:**
- head points to the node with value 4.
- Print 4 ->.
- Move head to nil (end of the list).
6. **End of the List:**
- Since head is now nil, the loop terminates.
- Print nil to indicate the end of the linked list.

### Leetcode 141 - Linked List Cycle

**Example 1: [3, 2, 0, -4], pos = 1**

1. **Create the List:**
- Nodes: 3 -> 2 -> 0 -> -4
- Create a cycle where -4 points back to 2 (position 1).
List with cycle: 3 -> 2 -> 0 -> -4 -> 2 -> ...
2. **Apply Floyd’s Algorithm:**
- slow starts at 3, fast starts at 3.
- After first iteration:
- - slow moves to 2
- - fast moves to 0
- After second iteration:
- - slow moves to 0
- - fast moves to -4
- After third iteration:
- - slow moves to -4
- - fast moves to 2
- After fourth iteration:
- - slow moves to 2
- - fast moves to 2 (they meet)
**Result:** Cycle detected (true).

### Leetcode 19 - Remove Nth Node from End of List

**Example: head = [1,2,3,4,5], n = 2**

1. **Create the List:**
- head is 1 -> 2 -> 3 -> 4 -> 5
2. **Initialize Pointers:**
- dummy points to the head of the list.
- fast and slow both start at dummy.
3. **Advance the Fast Pointer:**
- Move fast n + 1 steps ahead (in this case, 3 steps):
- - fast moves to 1 (head), then 2, and then 3.
4. **Move Both Pointers Simultaneously:**
- Move both fast and slow one step at a time until fast reaches the end of the list:
- - fast moves to 4, 5, and then nil.
- - slow moves to 2, 3, and then 4.
5. **Remove the Node:**
- slow now points to the node just before the node we want to remove (4 in this case). So, update slow.Next to skip the node 4:
- slow.Next should point to 5, effectively removing 4.
6. **Result:**
- The modified list is: 1 -> 2 -> 3 -> 5

**Function Definition**
go
Copy code
```
func createList(arr []int) *ListNode {
    if len(arr) == 0 {
        return nil
    }
    head := &ListNode{Val: arr[0]}
    current := head
    for _, val := range arr[1:] {
        current.Next = &ListNode{Val: val}
        current = current.Next
    }
    return head
}
```
**Code Explanation**
1. **Check for Empty Array:**
go
Copy code
```
if len(arr) == 0 {
    return nil
}
```
- If the input slice arr is empty, the function returns nil, indicating that there is no linked list to create. This is a guard clause to handle edge cases where no nodes should be created.
2. **Create the Head Node:**
go
Copy code
```
head := &ListNode{Val: arr[0]}
```
- This line creates the first node of the linked list (head) with the value of the first element in the array (arr[0]).
- head is a pointer to this newly created node.
3. **Initialize the current Pointer:**
go
Copy code
```
current := head
```
- current is a pointer initialized to head. It will be used to traverse and build the linked list.
4. **Build the Rest of the List:**
go
Copy code
```
for _, val := range arr[1:] {
    current.Next = &ListNode{Val: val}
    current = current.Next
}
```
- This loop iterates over the rest of the array (arr[1:]), starting from the second element.
- For each element val in the remaining array, it creates a new ListNode with Val set to val.
- The Next pointer of the current node is set to point to this new node.
- The current pointer is then updated to this newly created node, so that the next iteration will append a new node after the current one.
5. **Return the Head of the List:**
go
Copy code
```
return head
```
- After the loop completes, head points to the first node of the fully constructed linked list.
- The function returns head, which is the starting point of the linked list.

**Example Usage**

**Given an array arr:**
go
Copy code
```
arr := []int{1, 2, 3, 4, 5}
```
**The createList function will:**

1. Create the head node with value 1.
2. Create subsequent nodes with values 2, 3, 4, and 5.
3. Link these nodes together to form the list: 1 -> 2 -> 3 -> 4 -> 5.

### Problem 54 - Spiral Matrix

1. Initialize Boundaries and Result List

Boundaries:
top = 0
bottom = 2
left = 0
right = 2
Result List: [] (Initially empty)
2. Traverse the Matrix in Spiral Order

**Iteration 1: Move Right along the Top Row (Row 0)**

- **Condition:** top <= bottom (0 ≤ 2) and left <= right (0 ≤ 2)
- **Action:** Traverse from left (0) to right (2) on the top row (row 0).
- - **Traverse:**
- - - Add matrix[0][0] (1) to result
- - - Add matrix[0][1] (2) to result
- - - Add matrix[0][2] (3) to result
**Update:** Increment top from 0 to 1
**Result:** [1, 2, 3]

**Updated Boundaries:**

- top = 1
- bottom = 2
- left = 0
- right = 2

**Iteration 2: Move Down along the Right Column (Column 2)**

- **Condition:** top <= bottom (1 ≤ 2) and left <= right (0 ≤ 2)
- **Action:** Traverse from top (1) to bottom (2) on the right column (column 2).
- - **Traverse:**
- - - Add matrix[1][2] (6) to result
- - - Add matrix[2][2] (9) to result
**Update:** Decrement right from 2 to 1
**Result:** [1, 2, 3, 6, 9]

**Updated Boundaries:**

- top = 1
- bottom = 2
- left = 0
- right = 1

**Iteration 3: Move Left along the Bottom Row (Row 2)**

- Condition: top <= bottom (1 ≤ 2) and left <= right (0 ≤ 1)
- Action: Traverse from right (1) to left (0) on the bottom row (row 2).
- - Traverse:
- - - Add matrix[2][1] (8) to result
- - - Add matrix[2][0] (7) to result
- Update: Decrement bottom from 2 to 1
- Result: [1, 2, 3, 6, 9, 8, 7]

Updated Boundaries:

- top = 1
- bottom = 1
- left = 0
- right = 1

**Iteration 4: Move Up along the Left Column (Column 0)**

- **Condition:** top <= bottom (1 ≤ 1) and left <= right (0 ≤ 1)
- **Action:** Traverse from bottom (1) to top (1) on the left column (column 0).
- - **Traverse:**
- - - Add matrix[1][0] (4) to result
- **Update:** Increment left from 0 to 1
- **Result:** [1, 2, 3, 6, 9, 8, 7, 4]

Updated Boundaries:

- top = 1
- bottom = 1
- left = 1
- right = 1

**Iteration 5: Move Right along the Remaining Top Row (Row 1)**

- **Condition:** top <= bottom (1 ≤ 1) and left <= right (1 ≤ 1)
- **Action:** Traverse from left (1) to right (1) on the top row (row 1).
- **Traverse:**
- - Add matrix[1][1] (5) to result
- **Update:** Increment top from 1 to 2
**Result:** [1, 2, 3, 6, 9, 8, 7, 4, 5]

**Updated Boundaries:**

- top = 2
- bottom = 1
- left = 2
- right = 0
**Termination**
The loop stops because top > bottom and left > right are no longer valid conditions.

**Final Output**
The matrix elements in spiral order are: [1, 2, 3, 6, 9, 8, 7, 4, 5].


### Leetcode 242 - Valid Anagram

```
for _, char := range s {
    freq[char]++
}
```

**Iteration Details for s = "anagram":**

- Character 'a':
- - freq['a'] was 0, now becomes 1
- - freq = {'a': 1}

```
for _, char := range t {
    freq[char]--
    if freq[char] < 0 {
        return false
    }
}
```

**Iteration Details for t = "nagaram":**

- Character 'n':
- - freq['n'] was 1, now becomes 0
- - freq = {'a': 3, 'n': 0, 'g': 1, 'r': 1, 'm': 1}

