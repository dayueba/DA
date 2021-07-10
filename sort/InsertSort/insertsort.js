const insertSort = (arr) => {
    for (let i = 1; i < arr.length; i++) {
        const val = arr[i];
        let j = i;
        for (; j > 0 && val < arr[j-1]; j--) {
            arr[j] = arr[j-1];
        }
        arr[j] = val;
    }
    console.log(arr)
}

(function () {
    const arr = [4, 2, 3, 1, 9, 6, 7];
    insertSort(arr);
})();