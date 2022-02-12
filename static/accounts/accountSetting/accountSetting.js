// 头部导航栏部分
const noneAll = document.querySelectorAll('.nav-left a');
const movie = document.querySelector('.movie');
const settings = document.querySelector('#settings');
const myPage = document.querySelector('#my-page');
const none1 = document.querySelector('#none1');
const none2 = document.querySelector('#none2');
const accountSetting = document.querySelector('#account-setting');
const logout = document.querySelector('#logout');
document.querySelector('.account')['onclick'] = () => {
    if (settings.style.display == 'none') {
        settings.style.display = 'block';
    } else {
        settings.style.display = 'none';
    }
}
for (let i = 0; i < noneAll.length; i++) {
    noneAll[i].onclick = () => {
        alert('没有这个功能😭')
    }
}
movie.onclick = () => {
    window.location.href = "../../../index/main/main.html"
}
myPage.onclick = () => {

}
none1.onclick = () => {
    alert('没有这个功能😭')
}
none2.onclick = () => {
    alert('没有这个功能😭')
}
accountSetting.onclick = () => {
    window.location.href = "../accountSetting.html"
}
logout.onclick = () => {
    localStorage.clear();
    window.location.href = "../../../index/main/main.html"
}

// 点击切换功能
const otherSetup = document.querySelector('#other-setup');
const basicSetup = document.querySelector('#basic-setup');
const basic = document.querySelector('.basic');
const others = document.querySelector('.others');
otherSetup.onclick = () => {
    basicSetup.className = 'unselected';
    otherSetup.className = 'selected';
    basic.style.display = 'none';
    others.style.display = 'block';
}
basicSetup.onclick = () => {
    basicSetup.className = 'selected';
    otherSetup.className = 'unselected';
    basic.style.display = 'flex';
    others.style.display = 'none';
}