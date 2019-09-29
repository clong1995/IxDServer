cp.ready(() => {
    cp.on(".show", document, "click", t => {
        let infoDom = cp.query(".info", t.parentNode);
        cp.hasClass(infoDom, "hide") ? cp.show(infoDom) : cp.hide(infoDom)
    });

    //添加
    cp.on(".submit", 'click', t => {
        let parent = t.parentNode;
        let api = cp.query('.api', parent).value;
        let requestDom = cp.query('.request', parent);
        let data = requestDom ? JSON.parse(requestDom.value) : null;
        let responseDom = cp.query('.response', parent);
        cp.request(CONFIG.api + api, data, {
            headers: {
                Authorization: localStorage.getItem("Authorization")
            }
        }).then(res => {
            //responseDom.value = JSON.stringify(res);
            $(responseDom).JSONView(JSON.stringify(res));
            //登录和续期保存token
            if (api === ":19959/auth/signin" || api === ":19959/auth/prolong") {
                if (res.code === 0) {
                    //保存token
                    localStorage.setItem("Authorization", res.data)
                }
            }
        }).catch(err => {
            responseDom.value = JSON.stringify(err)
        })
    });
});