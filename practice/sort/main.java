public int[] mergeSort(int[] arr) {
    if (arr.length < 2) return arr;
    mergeSort(arr, 0, arr.length - 1);
    return arr;
}

private void mergeSort(int[] arr, int left, int right) {
    if(left < right) {
        int center = left + (right - left) / 2;
        mergeSort(arr, left, center);
        mergeSort(arr, center + 1, right);
        merge(arr, left, center, right);
    }
}

// 原地归并（手摇算法）
private void merge(int[] arr, int leftPos, int leftEnd, int rightEnd) {
    int i = leftPos, j = leftEnd + 1; // #1
    while(i < j && j <= rightEnd) {
        while(i < j && arr[i] <= arr[j]) i++; // #2
        int index = j; // #3
        while(j <= rightEnd && arr[j] < arr[i]) j++; // #4 注意是 arr[j] < arr[i]，即找到j使得arr[j] 为第一个大于等于 arr[i]值
        exchange(arr, i, index - 1, j - 1); // #5
    }
}

// 三次翻转实现交换
private void exchange(int[] arr, int left, int leftEnd, int rightEnd) {
    reverse(arr, left, leftEnd);
    reverse(arr, leftEnd + 1, rightEnd);
    reverse(arr, left, rightEnd);
}

private void reverse(int[] arr, int start, int end) {
    while(start < end) {
        swap(arr, start, end);
        start++;
        end--;
    }
}
