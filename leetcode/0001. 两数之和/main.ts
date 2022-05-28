function twoSum1(nums: number[], target: number): number[] {
    let map = new Map<number, number>();
    for (let i = 0; i < nums.length; i++) {
        let diff: number = target - nums[i];
        if (map.has(diff)) {
            return [map.get(diff) as number, i];
        }
        map.set(nums[i], i);
    }
    return [];
}

(function main() {
    console.log(twoSum1([1, 2, 3], 3));
})()
