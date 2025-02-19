### For problem 1 Two-Sum

Add later 

---

### For problem 3 Longest Substring Without Repeating Characters

Sure! Let’s dive deeper into the part `maxLength = max(maxLength, i - start + 1)` and how this relates to the sliding window concept, along with the idea of expanding and contracting the window when characters are repeated.

### **Sliding Window Concept:**

In the context of this problem, **sliding window** means you are maintaining a "window" (subarray or substring) that contains unique characters as you move through the string. The window starts as a small substring and expands by moving the `i` pointer (the right boundary of the window) as you process each character in the string.

When a character is repeated and already exists inside the window (i.e., inside the range `[start, i]`), you need to adjust the `start` pointer to shrink the window and remove the duplicate.

### **Understanding `i - start + 1`:**

1. **`i`** is the current index in the string, and `start` is the index of the beginning of the window.
2. `i - start` gives you the length of the window, **excluding** the character at `start`.
3. **Adding 1** accounts for the current character `i`, since the window is **inclusive** of both `start` and `i`.

So, `i - start + 1` calculates the **current length of the window** that holds unique characters. 

- **Why `+1`?**
  - Think of `i - start` as the difference between the two pointer indices, but this difference gives you the length excluding the current character at index `i`. To get the **actual length of the window**, you need to include that current character, so you add `1`.

### **Expanding the Window:**

When you process each character, if it hasn’t been seen before (or is outside the current window), you **expand the window** by including that character.

For example:
- Let's say the substring currently under consideration is `"abc"` and `i = 2` (meaning you are at the third character of the string). The window length is `i - start + 1`, so for this substring, it will be `2 - 0 + 1 = 3`.

In this case, `maxLength = max(maxLength, i - start + 1)` will be:
- Initially, `maxLength = 0` (before we start processing).
- After processing `"abc"`, `maxLength = max(0, 3)`, so `maxLength` becomes 3.

This means we've updated the maximum length to be the length of this substring without repeating characters.

### **Contracting the Window (When a Character is Repeated):**

When you encounter a repeated character **within the current window**, you need to **shrink the window** from the left to make sure all characters in the window remain unique. This is where the `start` pointer comes into play.

For example:
- Let's say the current string is `"abcabcbb"`, and you're at the 4th character, `s[4] = 'a'`.
- The character `'a'` was already seen at `start = 0`, so the window from `start = 0` to `i = 4` is invalid because it contains a repeated `'a'`.

To make the window valid again, you move the `start` pointer to **just after** the last occurrence of `'a'` (which is index 1). So, `start` will move from 0 to 1, and now the window will be `"bca"`.

The length of this valid window will now be `i - start + 1 = 4 - 1 + 1 = 4`. You continue this process until there are no duplicates in the window, and the `maxLength` is updated accordingly.

### **Let’s Walk Through an Example:**

Consider the input string `"abcabcbb"`.

1. **First window (`start = 0`, `i = 0`)**:  
   Window = `"a"`, length = `0 - 0 + 1 = 1`.  
   `maxLength` = `max(0, 1) = 1`.

2. **Second window (`start = 0`, `i = 1`)**:  
   Window = `"ab"`, length = `1 - 0 + 1 = 2`.  
   `maxLength` = `max(1, 2) = 2`.

3. **Third window (`start = 0`, `i = 2`)**:  
   Window = `"abc"`, length = `2 - 0 + 1 = 3`.  
   `maxLength` = `max(2, 3) = 3`.

4. **Fourth window (`start = 0`, `i = 3`)**:  
   Window = `"abca"`, but since `'a'` is repeating (seen at index 0), we move `start` to `start = 1` (just after the last occurrence of `'a'`).  
   Now the window = `"bca"`, length = `3 - 1 + 1 = 3`.  
   `maxLength` remains `3`.

5. **Fifth window (`start = 1`, `i = 4`)**:  
   Window = `"bcab"`, but since `'b'` is repeating (seen at index 1), we move `start` to `start = 2` (just after the last occurrence of `'b'`).  
   Now the window = `"cab"`, length = `4 - 2 + 1 = 3`.  
   `maxLength` remains `3`.

6. **Sixth window (`start = 2`, `i = 5`)**:  
   Window = `"cabc"`, but since `'c'` is repeating (seen at index 2), we move `start` to `start = 3` (just after the last occurrence of `'c'`).  
   Now the window = `"abc"`, length = `5 - 3 + 1 = 3`.  
   `maxLength` remains `3`.

7. **Seventh window (`start = 3`, `i = 6`)**:  
   Window = `"abcb"`, but since `'b'` is repeating (seen at index 4), we move `start` to `start = 5` (just after the last occurrence of `'b'`).  
   Now the window = `"cb"`, length = `6 - 5 + 1 = 2`.  
   `maxLength` remains `3`.

8. **Eighth window (`start = 5`, `i = 7`)**:  
   Window = `"cbb"`, but since `'b'` is repeating (seen at index 6), we move `start` to `start = 7` (just after the last occurrence of `'b'`).  
   Now the window = `"b"`, length = `7 - 7 + 1 = 1`.  
   `maxLength` remains `3`.

### **Key Points:**

- **Expanding** the window means you keep extending the current substring by moving `i` forward and adding new unique characters.
- **Contracting** the window happens when you encounter a duplicate character inside the window. You move `start` to the right, just after the last occurrence of the repeated character, to ensure the window only contains unique characters.
- **`i - start + 1`** calculates the length of the current valid window.

### **In summary**:
- The **sliding window** allows you to dynamically manage the substring of unique characters.
- **Expanding** the window happens when you add new characters.
- **Contracting** happens when you remove duplicates by adjusting the `start` pointer.



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

---

### Problem 11 Container With Most Water

Add later


---

### For problem 53 Maximum Subarray
Solving using Kadane's Algorithm.

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

---

### For Problem 152 

#### Key Idea:
We are trying to find the maximum product of any contiguous subarray in the array. The challenge arises because multiplying negative numbers can flip the sign of the product. So, it's essential to track both the **maximum** and **minimum** product up to the current index.

#### Why Track `maxProd` and `minProd`?
The reason we keep track of both `maxProd` (the maximum product so far) and `minProd` (the minimum product so far) is that a negative number can cause a small negative product (minProd) to turn into a larger positive product when multiplied. Similarly, the maximum product may decrease if we multiply by a negative number.

For example:

- In an array like `[2, 3, -2, 4]`, the negative number `-2` can turn a **small negative** product into a **large positive** product when multiplied by another negative number, making it necessary to track both the largest and smallest products.

#### Explanation of `maxProd = max(nums[i], maxProd * nums[i])` and `minProd = min(nums[i], minProd * nums[i])`:
1. `maxProd = max(nums[i], maxProd * nums[i])`:
- `nums[i]`: This considers the case where we start a new subarray at index `i`. For example, when `nums[i]` is large or when the product up to the previous index is negative (we reset the subarray).
- `maxProd * nums[i]`: This considers the case where we continue the subarray by multiplying the current element with the previous product. If the previous product was positive, it would likely increase the current product. If it was negative, the product might increase (if multiplied by a negative number).

**Example**: Consider the array `[2, 3, -2, 4]`:

- After processing `2` and `3`, the `maxProd` becomes `6` (from `2 * 3`).
- But when we encounter `-2`, we might have two possibilities:
  - `maxProd * (-2)` results in a **negative** value (not ideal for max).
  - Start fresh with `nums[i] = -2` (this gives a potential smaller number that can later be turned into a positive product).
2. `minProd = min(nums[i], minProd * nums[i])`:
- `nums[i]`: This considers the case where we start a new subarray at index `i` with just the current element, which might be a negative number or a small value.
- `minProd * nums[i]`: This considers the case where we extend the subarray by multiplying the current number with the previous minimum product. If the previous minimum product was negative, multiplying by a negative number could make it positive.

**Example**: Continuing from the previous example, if the `minProd` at a certain point is `-6` and the next number is `-2`, multiplying `-6 * -2` would result in a positive number. So, the minimum value might be important because it could help flip negative results to positive ones.

The following example is to understand when a grater number comes how the code perform iteration's in it.
#### Step-by-Step Walkthrough of the Example `[2, 3, -2, 8, 4]`:
Let's consider how this works for the array `[2, 3, -2, 8, 4]`:

- **Initialization**:
  - maxProd = 2 (initial value from nums[0])
  - minProd = 2
  - result = 2
- **Iteration 1** (i = 1, nums[1] = 3):
  - maxProd = max(3, 2 * 3) = 6
  - minProd = min(3, 2 * 3) = 3
  - result = max(2, 6) = 6
- **Iteration 2** (i = 2, nums[2] = -2):
  - maxProd = max(-2, 6 * -2) = 12 (flip to positive)
  - minProd = min(-2, 3 * -2) = -6 (becomes negative)
  - result = max(6, 12) = 12
- **Iteration 3** (i = 3, nums[3] = 8):
  - maxProd = max(8, 12 * 8) = 96
  - minProd = min(8, -6 * 8) = -48
  - result = max(12, 96) = 96
- **Iteration 4** (i = 4, nums[4] = 4):
  - maxProd = max(4, 96 * 4) = 384
  - minProd = min(4, -48 * 4) = -192
  - result = max(96, 384) = 384

**Result**:
At the end of the loop, `result` will be `384`, which is the maximum product of any contiguous subarray.

The following example is to understand when two negative number are there how they will become positive.

#### Example Array: `[2, -3, -2, 4]`

- **Initialization**:
  - maxProd = 2 (initial value from nums[0])
  - minProd = 2
  - result = 2
- **Iteration 1** (i = 1, nums[1] = -3):

**Handle negative number**:

Since `nums[1] = -3` is negative, we swap `maxProd` and `minProd`:

- `minProd = 2` (previous `maxProd`)
- `maxProd = 2` (previous `minProd`)

  - maxProd = max(-3, 2 * -3) = -6
  - minProd = min(3, 2 * -3) = -3
  - result = max(2, -6) = 2
- **Iteration 2** (i = 2, nums[2] = -2):

**Handle negative number**:

Since `nums[1] = -2` is negative, we swap `maxProd` and `minProd`:

- `minProd = -3` (previous `maxProd`)
- `maxProd = -6` (previous `minProd`)

  - maxProd = max(-2, -6 * -2) = 12 (flip to positive)
  - minProd = min(-2, -3 * -2) = -6 (becomes negative)
  - result = max(6, 12) = 12
- **Iteration 3** (i = 3, nums[3] = 8):
  - maxProd = max(4, 12 * 4) = 48
  - minProd = min(4, -6 * 4) = -24
  - result = max(12, 48) = 48

**Result**:
At the end of the loop, `result` will be `48`, which is the maximum product of any contiguous subarray.

**Key Points**:
- When a negative number appears, we swap the `maxProd` and `minProd` because multiplying a negative number with a negative product can turn it into a large positive product.
- The final result is the `maximum product` that was encountered at any point in the array, which is `48` in this example.

---

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

---

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

---

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

---

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

---

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

---

### For Problem 424 - Longest Repeating Character Replacement

I understand your confusion. Let's break it down and walk through the example in more detail. I'll explain the working of the sliding window, `charCount[s[right]]++`, how `maxFreq` works, and why we need `windowSize`.

---

### Goal:
We want to find the **longest substring** that can be obtained by replacing **at most `k` characters** to make all characters in that substring identical. We will use a sliding window approach to efficiently solve this problem.

### Key Variables:
1. **`charCount`:** A map to keep track of the frequency of characters in the current window.
2. **`maxFreq`:** The maximum frequency of a character in the current window. We use this to check if the window is valid (if we need to shrink it).
3. **`windowSize`:** This is the current size of the window, which is calculated as `right - left + 1`.
4. **`maxLength`:** The longest valid window (substring) found so far.

---

### How `charCount[s[right]]++` Works:
Whenever we encounter a character at index `right` (the right boundary of our window), we increase its count in the `charCount` map. This helps us keep track of the frequency of characters in the current window.

For example, if we encounter `'A'` at index `right = 0`, it will update the count in `charCount`.

---

### Why `maxFreq` is Important:
`maxFreq` tracks the frequency of the most frequent character in the current window. We need to track this because, to make a substring of all identical characters by replacing at most `k` characters, we need to know how many characters in the current window are already the same as the most frequent character.

If the **current window size (`right - left + 1`)** minus `maxFreq` is **greater than `k`**, it means we have more than `k` characters that are different from the most frequent character. In this case, we must shrink the window by moving the `left` pointer to the right to make it valid again.

---

### Why We Need to Calculate `windowSize`:
`windowSize` is simply the number of characters in the current window, i.e., `right - left + 1`. This helps us determine how many characters we have in the window and whether it's valid or not.

For the window to be valid (i.e., we can replace at most `k` characters to make the substring uniform), the number of characters to be replaced should be less than or equal to `k`. The number of characters to be replaced is:
```
windowSize - maxFreq
```
If this value is **greater than `k`**, we shrink the window from the left.

---

### Conclusion:
- **How `charCount[s[right]]++` works**: This keeps track of the frequency of characters in the current window, helping to determine the most frequent character (`maxFreq`).
- **Why `maxFreq` helps**: It allows us to know the number of characters we need to replace. If the characters to replace exceed `k`, we shrink the window.
- **Why `windowSize` is important**: It helps us determine if the current window is valid. If the number of characters to replace exceeds `k`, we shrink the window.

The sliding window adjusts based on the need to maintain a valid window while iterating through the string.

---

Great question! Let's break this down carefully so that you understand how it works and why `'A'` is subtracted.

### Why do we subtract `'A'` from `s[wEnd]` and `s[wStart]`?

In the problem you're working on, you're dealing with a string of uppercase English letters, and you need to count the occurrences of each character in a sliding window. Here's the key point: 

1. **We have 26 possible characters** (from 'A' to 'Z') in the string.
2. **We want to map these characters to indices** in an array to count how often each character appears in the window.

#### Why `s[wEnd] - 'A'` works:

- In Go, when you subtract a character from another character, you're subtracting their **ASCII values**.
- The ASCII value of `'A'` is 65, `'B'` is 66, `'C'` is 67, and so on. The characters are mapped to integer values where `'A'` corresponds to `0`, `'B'` corresponds to `1`, `'C'` corresponds to `2`, and so on.
  
- **`s[wEnd] - 'A'`** is a way of converting the character into an index of the `freqs` array. For example:
  - If `s[wEnd]` is `'A'`, then `s[wEnd] - 'A' = 0` → This corresponds to the 0th index of the `freqs` array.
  - If `s[wEnd]` is `'B'`, then `s[wEnd] - 'A' = 1` → This corresponds to the 1st index of the `freqs` array.
  - If `s[wEnd]` is `'C'`, then `s[wEnd] - 'A' = 2` → This corresponds to the 2nd index of the `freqs` array.
  - And so on, up to `'Z'`, where `s[wEnd] - 'A' = 25`.

### Why `freqs[s[wStart] - 'A']--`?

Similarly, when you're shrinking the window by moving `wStart` to the right (removing the character at `s[wStart]` from the window), you need to **decrease the frequency of the character** that `wStart` points to.

- **`s[wStart] - 'A'`** gives the **index** of the character in the `freqs` array. This is the **same reasoning** as `s[wEnd] - 'A'`. We subtract `'A'` to convert the character into an integer index, so we know where to decrease the frequency in the array.
- For example:
  - If `s[wStart]` is `'A'`, then `s[wStart] - 'A' = 0` → It means we will decrease the count of the character `'A'` in the `freqs` array at index `0`.
  - If `s[wStart]` is `'B'`, then `s[wStart] - 'A' = 1` → It means we will decrease the count of the character `'B'` in the `freqs` array at index `1`.

### Here's how the sliding window and frequency count works:

1. You initialize an array `freqs` of length 26, with all elements set to 0. This array will store the frequency of each character in the current sliding window.

   ```go
   freqs := make([]int, 26)
   ```

   Initially, `freqs` will look like this:
   ```
   freqs = [0, 0, 0, ..., 0]  // 26 zeros (representing the frequencies of 'A' to 'Z')
   ```

2. As you expand the window (`wEnd` moves to the right), you update the frequency of the character at `s[wEnd]`.

   ```go
   ch := s[wEnd] - 'A'
   freqs[ch]++  // Increment the count of the character at s[wEnd]
   ```

   If `s[wEnd] = 'A'`, then `ch = 'A' - 'A' = 0`, so `freqs[0]++` will increase the count for 'A'.
   
   If `s[wEnd] = 'B'`, then `ch = 'B' - 'A' = 1`, so `freqs[1]++` will increase the count for 'B'.

3. If the window is **invalid** (i.e., you need more than `k` replacements to make it valid), you shrink the window from the left by moving `wStart` to the right.

   ```go
   freqs[s[wStart] - 'A']--  // Decrease the count of the character at s[wStart]
   wStart++  // Move the left boundary of the window to the right
   ```

   If `s[wStart] = 'A'`, then `s[wStart] - 'A' = 0`, so `freqs[0]--` will decrease the count for 'A' as you're removing that character from the window.

4. The **maximum frequency** (`maximum`) is updated each time you expand the window, and it represents the most frequent character in the current window. The idea is that if the window size minus the frequency of the most frequent character exceeds `k`, you need to shrink the window.

---

### Example: How it works with `"AABABBA"`

Let's walk through an example with the string `s = "AABABBA"` and `k = 1`.

1. **Initialize variables:**

   ```
   freqs = [0, 0, 0, ..., 0]   // Frequency of characters 'A' to 'Z'
   wStart = 0
   maximum = 0
   result = 0
   ```

2. **Iteration 1:**
   - `wEnd = 0`, `s[wEnd] = 'A'`.
   - `ch = 'A' - 'A' = 0`, so `freqs[0]++` → `freqs = [1, 0, 0, ..., 0]`.
   - `maximum = max(0, 1)` → `maximum = 1`.
   - Window size: `wEnd - wStart + 1 = 1`.
   - Window is valid (`1 - 1 <= 1`), update `result = max(0, 1)` → `result = 1`.

3. **Iteration 2:**
   - `wEnd = 1`, `s[wEnd] = 'A'`.
   - `ch = 'A' - 'A' = 0`, so `freqs[0]++` → `freqs = [2, 0, 0, ..., 0]`.
   - `maximum = max(1, 2)` → `maximum = 2`.
   - Window size: `wEnd - wStart + 1 = 2`.
   - Window is valid (`2 - 2 <= 1`), update `result = max(1, 2)` → `result = 2`.

4. **Iteration 3:**
   - `wEnd = 2`, `s[wEnd] = 'B'`.
   - `ch = 'B' - 'A' = 1`, so `freqs[1]++` → `freqs = [2, 1, 0, ..., 0]`.
   - `maximum = max(2, 1)` → `maximum = 2`.
   - Window size: `wEnd - wStart + 1 = 3`.
   - Window is valid (`3 - 2 <= 1`), update `result = max(2, 3)` → `result = 3`.

5. **Iteration 4:**
   - `wEnd = 3`, `s[wEnd] = 'A'`.
   - `ch = 'A' - 'A' = 0`, so `freqs[0]++` → `freqs = [3, 1, 0, ..., 0]`.
   - `maximum = max(2, 3)` → `maximum = 3`.
   - Window size: `wEnd - wStart + 1 = 4`.
   - Window is valid (`4 - 3 <= 1`), update `result = max(3, 4)` → `result = 4`.

6. **Iteration 5:**
   - `wEnd = 4`, `s[wEnd] = 'B'`.
   - `ch = 'B' - 'A' = 1`, so `freqs[1]++` → `freqs = [3, 2, 0, ..., 0]`.
   - `maximum = max(3, 2)` → `maximum = 3`.
   - Window size: `wEnd - wStart + 1 = 5`.
   - Window is valid (`5 - 3 <= 1`), update `result = max(4, 5)` → `result = 5`.  ChatGpt has done a mistake here as 5 - 3 is 2 not 1. 

---

### Summary:
- **Subtracting `'A'`:** Converts characters into zero-based indices (i.e., `'A' = 0, 'B' = 1, ..., 'Z' = 25`) to index into the `freqs` array.
- **`freqs[s[wStart] - 'A']--`:** Decreases the frequency of the character at the `wStart` index in the array when the window shrinks.
- **Sliding window logic:** Expands the window by moving `wEnd`, and shrinks the window by

 moving `wStart` to keep the window valid.

This technique efficiently counts the characters and keeps track of the most frequent character to determine if the window is valid based on the allowed number of replacements `k`.

---

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

---

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

---

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

---


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

---

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

---

Yes, you are absolutely right in pointing out that we are comparing `sFreq[key]` with the value from `tFreq`, not directly checking `tFreq['a']`. Let me clarify the logic more clearly.

### What We Are Doing:
In the loop:

```go
for key, value := range tFreq {
    if sFreq[key] != value {
        return false
    }
}
```

Here, the `key` represents the character, and the `value` represents the frequency of that character in the `t` string (`tFreq` map). We are checking if the **frequency of the character in `sFreq` is the same as the frequency of the same character in `tFreq`**.

### Breakdown:
- `key`: This is a character from `tFreq` (e.g., `'a'`, `'b'`, etc.).
- `value`: This is the frequency of that character in string `t`.
- `sFreq[key]`: This is the frequency of that **same character** in string `s`.

---

Certainly! Let's break down and visualize the part where you are updating the counts for the characters in both strings `s` and `t`:

### Code:
```go
for i := 0; i < len(s); i++ {
    charCount[s[i]]++  // Increment the count for the character in s
    charCount[t[i]]--  // Decrement the count for the character in t
}

// Check if all counts are zero
for _, count := range charCount {
    if count != 0 {
        return false  // If any count is not zero, return false
    }
}
```



### Explanation For ChatGPT Solution:

- **`charCount[s[i]]++`**: For each character in `s`, you increment the corresponding count in the map `charCount`.
- **`charCount[t[i]]--`**: For each character in `t`, you decrement the corresponding count in the map `charCount`.
- **Goal**: After processing both strings `s` and `t`, if `s` is an anagram of `t`, all counts should be zero. If they are not, it means the two strings do not have the same characters and frequencies, so the function returns `false`.

### Visualization:

Let's take an example to walk through it:  
- `s = "aabbcc"`  
- `t = "abcbac"`

### Initial Setup:

- Create a map `charCount` to track the character counts.

```go
charCount := make(map[byte]int)
```

At the start, `charCount` is empty.

### Iteration 1: i = 0 (First characters of `s` and `t`)
- `s[0] = 'a'`, `t[0] = 'a'`
- We do the following:
  - `charCount['a']++`: `charCount['a']` becomes 1.
  - `charCount['a']--`: `charCount['a']` becomes 0.

So, after the first iteration, `charCount` looks like:
```go
charCount = {'a': 0}
```

### Iteration 2: i = 1 (Second characters of `s` and `t`)
- `s[1] = 'a'`, `t[1] = 'b'`
- We do the following:
  - `charCount['a']++`: `charCount['a']` becomes 1.
  - `charCount['b']--`: `charCount['b']` becomes -1.

So, after the second iteration, `charCount` looks like:
```go
charCount = {'a': 1, 'b': -1}
```

### Iteration 3: i = 2 (Third characters of `s` and `t`)
- `s[2] = 'b'`, `t[2] = 'c'`
- We do the following:
  - `charCount['b']++`: `charCount['b']` becomes 0.
  - `charCount['c']--`: `charCount['c']` becomes -1.

So, after the third iteration, `charCount` looks like:
```go
charCount = {'a': 1, 'b': 0, 'c': -1}
```

---

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

---
