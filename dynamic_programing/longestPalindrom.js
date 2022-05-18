/**
 * @param {string} s
 * @return {string}
 */
 var longestPalindrome = function(s) {
    if(s == null || s.length < 1 ) {
        return '';
    }
    let start = 0;
    let end = 0;
    for(let i = 0; i<s.length; i++) {
        let c1 = palindromHength(s, i, i);
        let c2 = palindromHength(s, i, i + 1);
        let length = Math.max(c1, c2);
        if(length > end - start) {
            start = i - Math.floor((length - 1)/2);
            end = i + Math.floor(length/2);
        }
    }
    return s.substring(start, end + 1);
}

let palindromHength = (s, start, end) => {
    while(start>=0 && end<s.length && s.charAt(start) == s.charAt(end)){
        start --;
        end ++;
    }
    return end - start - 1
}