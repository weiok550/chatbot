<template>
  <div class="container">
    <div class="row">
      <div class="col-12 chatbox">
        <div class="col-12 col-lg-7 col-xl-9 d-flex flex-column">
          <div class="card-body scrollable" style="height: 35rem">
            <div class="chat">
              <div class="chat-bubbles">
                <div class="chat-item">
                  <div class="row align-items-end justify-content-end">
                    <div class="col col-lg-6">
                      <div class="chat-bubble chat-bubble-me">
                        <div class="chat-bubble-title">
                          <div class="row">
                            <div class="col chat-bubble-author">Paweł Kuna</div>
                            <div class="col-auto chat-bubble-date">09:32</div>
                          </div>
                        </div>
                        <div class="chat-bubble-body">
                          <p>
                            Hey guys, I just pushed a new commit on the
                            <code>dev</code> branch. Can you have a look and
                            tell me what you think?
                          </p>
                        </div>
                      </div>
                    </div>
                    <div class="col-auto">
                      <span
                        class="avatar"
                        style="background-image: url(./static/avatars/000m.jpg)"
                      ></span>
                    </div>
                  </div>
                </div>
                <div class="chat-item">
                  <div class="row align-items-end">
                    <div class="col-auto">
                      <span
                        class="avatar"
                        style="background-image: url(./static/avatars/002m.jpg)"
                      ></span>
                    </div>
                    <div class="col col-lg-6">
                      <div class="chat-bubble">
                        <div class="chat-bubble-title">
                          <div class="row">
                            <div class="col chat-bubble-author">
                              Mallory Hulme
                            </div>
                            <div class="col-auto chat-bubble-date">09:34</div>
                          </div>
                        </div>
                        <div class="chat-bubble-body">
                          <p>Sure Paweł, let me pull the latest updates.</p>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                
              </div>
            </div>
          </div>
          <div class="card-footer">
            <div class="input-group input-group-flat">
              <input
                type="text"
                class="form-control"
                autocomplete="off"
                placeholder="Type message"
              />
              <span class="input-group-text">
                <a
                  href="#"
                  class="link-secondary"
                  data-bs-toggle="tooltip"
                  aria-label="Clear search"
                  data-bs-original-title="Clear search"
                >
                  <!-- Download SVG icon from http://tabler-icons.io/i/mood-smile -->
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="icon"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    stroke-width="2"
                    stroke="currentColor"
                    fill="none"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                    <path d="M12 12m-9 0a9 9 0 1 0 18 0a9 9 0 1 0 -18 0"></path>
                    <path d="M9 10l.01 0"></path>
                    <path d="M15 10l.01 0"></path>
                    <path d="M9.5 15a3.5 3.5 0 0 0 5 0"></path>
                  </svg>
                </a>
                <a
                  href="#"
                  class="link-secondary ms-2"
                  data-bs-toggle="tooltip"
                  aria-label="Add notification"
                  data-bs-original-title="Add notification"
                >
                  <!-- Download SVG icon from http://tabler-icons.io/i/paperclip -->
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="icon"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    stroke-width="2"
                    stroke="currentColor"
                    fill="none"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                    <path
                      d="M15 7l-6.5 6.5a1.5 1.5 0 0 0 3 3l6.5 -6.5a3 3 0 0 0 -6 -6l-6.5 6.5a4.5 4.5 0 0 0 9 9l6.5 -6.5"
                    ></path>
                  </svg>
                </a>
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ChatPage",
};
</script>

<style scoped>
@import "../assets/css/chat-page.css";
</style>

<style scoped>
body {
  display: block;
  background-color: black;
}

.chatbox {
  background-color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
</style>

// 页面初始化请求接口 http://chatbot.putianxia.top:8080/api/chat/getChatRecords 请求参数为 start=0, 接口返回数据结构如下：
//   {
//     "code": 200,
//     "msg": "ok",
//     "data": {
//         "list": [
//             {
//                 "id": 10006,
//                 "type": 1,
//                 "user_id": 1,
//                 "msg": "你好",
//                 "created_at": 1710483288210
//             },
//             {
//                 "id": 10006,
//                 "type": 0,
//                 "user_id": 1,
//                 "msg": "你好",
//                 "created_at": 1710483288210
//             }
//         ],
//         "index" :1710483288210 
//       }
//   }
//   帮我实现如下功能：
//   1.  
//   1. 倒序遍历data.list将数组元素数据对应渲染到.chat-item 的div，如果元素type为1那么.align-items-end标签类后面同时有justify-content-end类否则没有，如果type为1 .chat-bubble-author标签的文本内容为"我" 否则为"机器人"，要把created_at 毫秒时间戳转为 mm-dd hh:ii:ss 格式放入.chat-bubble-date标签的文本内 容，把数组元素msg内容放入.chat-bubble-body标签下的p标签内；
//   2. 拿到接口数据记录data.index, 这个数据作为start参数 当向上滚动滚轮到card-body顶部时继续请求接口，传的start值就是这个data.index