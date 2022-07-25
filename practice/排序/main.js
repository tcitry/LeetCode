let l = [
  1, 2, 3, 4, 3, 2, 2, 3, 4, 45, 6, 5, 4, 3, 45, 5, 6, 5, 4, 33, 2, 4, 5,
];


let nums = [1, 2, 4, 3, 5, 6, 7, 8, 9, 10];

let nums1 = new Array(10, 2, 4, 3, 5, 6, 7, 8, 9, 100);

function selectSort(nums) {
  for (let i = 0; i < nums.length; i++) {
    let min = i;
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[j] < nums[min]) {
        min = j;
      }
    }
    if (min !== i) {
      let temp = nums[i];
      nums[i] = nums[min];
      nums[min] = temp;
    }
  }
  return nums;
}

function insertionSort(nums) {
  for (let i = 1; i < nums.length; i++) {
    let temp = nums[i];
    let j = i - 1;
    while (j >= 0 && nums[j] > temp) {
      nums[j + 1] = nums[j];
      j--;
    }
    nums[j + 1] = temp;
  }
  return nums;
}

function countSort(nums) {
  let min = Infinity;
  let max = -Infinity;

  for (let i = 0; i < nums.length; i++) {
    min = Math.min(min, nums[i]);
    max = Math.max(max, nums[i]);
  }

  let counts = new Array(nums.length).fill(0);
  for (let i = 0; i < nums.length; i++) {
    counts[nums[i] - min]++;
  }
  let i = 0;
  for (let num = min; num <= max; num++) {
    while (counts[num - min] > 0) {
      nums[i++] = num;
      counts[num - min]--;
    }
  }
  return nums;
}

function popoSort(nums) {
  for (let i = 0; i < nums.length; i++) {
    for (let j = 0; j < nums.length - i - 1; j++) {
      if (nums[j] > nums[j + 1]) {
        let temp = nums[j];
        nums[j] = nums[j + 1];
        nums[j + 1] = temp;
      }
    }
  }
  return nums;
}

function shellSort(nums) {
  let gap = Math.floor(nums.length / 2);
  while (gap > 0) {
    for (let i = gap; i < nums.length; i++) {
      let temp = nums[i];
      let j = i - gap;
      while (j >= 0 && nums[j] > temp) {
        nums[j + gap] = nums[j];
        j -= gap;
      }
      nums[j + gap] = temp;
    }
    gap = Math.floor(gap / 2);
  }
  return nums;
}

function heapSort(nums) {
  function heapify(nums, start, end) {
    let root = start;
    while (2 * root + 1 <= end) {
      let child = 2 * root + 1;
      let swap = root;
      if (nums[swap] < nums[child]) {
        swap = child;
      }
      if (child + 1 <= end && nums[swap] < nums[child + 1]) {
        swap = child + 1;
      }
      if (swap === root) {
        break;
      }
      let temp = nums[root];
      nums[root] = nums[swap];
      nums[swap] = temp;
      root = swap;
    }
  }
  for (let i = Math.floor(nums.length / 2); i >= 0; i--) {
    heapify(nums, i, nums.length - 1);
  }
  for (let i = nums.length - 1; i > 0; i--) {
    let temp = nums[0];
    nums[0] = nums[i];
    nums[i] = temp;
    heapify(nums, 0, i - 1);
  }
  return nums;
}

function mergeSort0(nums) {
  function merge(left, right) {
    let result = [];
    while (left.length && right.length) {
      if (left[0] < right[0]) {
        result.push(left.shift());
      } else {
        result.push(right.shift());
      }
    }
    while (left.length) {
      result.push(left.shift());
    }
    while (right.length) {
      result.push(right.shift());
    }
    return result;
  }
  if (nums.length < 2) {
    return nums;
  }
  let middle = Math.floor(nums.length / 2);
  let left = mergeSort(nums.slice(0, middle));
  let right = mergeSort(nums.slice(middle));
  return merge(left, right);
}

function quickSort(nums) {
  function partition(nums, start, end) {
    let pivot = nums[end];
    let i = start;
    for (let j = start; j < end; j++) {
      if (nums[j] <= pivot) {
        let temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
        i++;
      }
    }
    let temp = nums[i];
    nums[i] = nums[end];
    nums[end] = temp;
    return i;
  }
  function sort(nums, start, end) {
    if (start >= end) {
      return;
    }
    let pivot = partition(nums, start, end);
    sort(nums, start, pivot - 1);
    sort(nums, pivot + 1, end);
  }
  sort(nums, 0, nums.length - 1);
  return nums;
}


function quickSortStack(nums) {
  function sort(nums, start, end) {
    if (start >= end) {
      return;
    }
    let stack = [[start, end]];
    while (stack.length) {
      let [left, right] = stack.pop();
      if (left >= right) {
        continue;
      }
      let pivot = partition(nums, left, right);
      stack.push([left, pivot - 1]);
      stack.push([pivot + 1, right]);
    }
  }
  function partition(nums, start, end) {
    let pivot = nums[end];
    let i = start;
    for (let j = start; j < end; j++) {
      if (nums[j] <= pivot) {
        let temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
        i++;
      }
    }
    let temp = nums[i];
    nums[i] = nums[end];
    nums[end] = temp;
    return i;
  }
  sort(nums, 0, nums.length - 1);
  return nums;
}

function bucketSort(nums) {
  let max = Math.max(...nums);
  let min = Math.min(...nums);
  let bucket = new Array(max - min + 1).fill(0);
  for (let i = 0; i < nums.length; i++) {
    bucket[nums[i] - min]++;
  }
  let i = 0;
  for (let num = min; num <= max; num++) {
    while (bucket[num - min] > 0) {
      nums[i++] = num;
      bucket[num - min]--;
    }
  }
  return nums;
}

function radixSort(nums) {
  let max = Math.max(...nums);
  let radix = Math.floor(Math.log10(max)) + 1;
  for (let i = 0; i < radix; i++) {
    let bucket = new Array(10).fill(0);
    for (let j = 0; j < nums.length; j++) {
      let digit = Math.floor((nums[j] % Math.pow(10, i + 1)) / Math.pow(10, i));
      bucket[digit]++;
    }
    let i = 0;
    for (let num = 0; num < 10; num++) {
      while (bucket[num] > 0) {
        nums[i++] = num;
        bucket[num]--;
      }
    }
  }
  return nums;
}

console.log(quickSort(nums1));
