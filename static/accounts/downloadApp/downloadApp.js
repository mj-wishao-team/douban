const goToTop = document.querySelector('#go-to-top');
// 开始滚动时出现回顶部按钮
document.addEventListener('scroll', () => {
    if (window.pageYOffset >= 120){
        goToTop.style.display = 'block'
    } else {
        goToTop.style.display = 'none'
    }
});
// 写个小动画函数
function animate(obj, target, time) {
    let timer;
    clearInterval(obj.timer);
    obj.timer = setInterval( () => {
        let step = (target - window.pageYOffset) / 10;   
        step = step > 0 ? Math.ceil(step) : Math.floor(step);
        if (window.pageYOffset == target) {
            clearInterval(obj.timer);
        }else {
            window.scroll(0,window.pageYOffset + step)
        }
    }, time);
};
// 实现点击滚回顶部效果
goToTop.onclick = () => {
    animate(window, 0, 15);
};
// 实现字在滚动到此处时的过渡动画
// 先鸽了，有时间再想想怎么做 TwT
