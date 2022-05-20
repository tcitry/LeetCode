var reverse = function (x) {
    let flag = x > 0 ? 1 : -1;
    let num = Math.abs(x).toString().split('').reverse().join('');
    if (parseInt(num) > Math.pow(2, 31) - 1) {
        return 0;
    }
    return flag * parseInt(num);
};