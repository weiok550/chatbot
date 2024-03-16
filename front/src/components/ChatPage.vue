<template>
  <div class="container">
    <div class="row">
      <div class="col-12 chatbox">
        <div class="col-12 col-lg-7 col-xl-9 d-flex flex-column">
          <div class="card-body scrollable" style="height: 49rem">
            <div class="chat">
              <div class="chat-bubbles">
                <!--历史聊天数据渲染 start-->
                <div
                  class="chat-item"
                  v-for="item in data.list.slice().reverse()"
                  :key="item.id"
                >
                  <div
                    :class="[
                      'row',
                      item.type === 1
                        ? 'align-items-end justify-content-end'
                        : '',
                    ]"
                  >
                    <div class="col col-lg-9">
                      <div
                        class="chat-bubble"
                        :class="item.type === 1 ? 'chat-bubble-me' : ''"
                      >
                        <div class="chat-bubble-title">
                          <div v-if="item.type === 0">
                            <div class="row">
                              <div class="col chat-bubble-author">机器人</div>
                              <div class="col-auto chat-bubble-date">
                                {{ formatDate(item.created_at) }}
                              </div>
                            </div>
                          </div>
                          <div v-else-if="item.type === 1">
                            <div class="row">
                              <div class="col chat-bubble-date">
                                {{ formatDate(item.created_at) }}
                              </div>
                              <div class="col-auto chat-bubble-author">
                                自己
                              </div>
                            </div>
                          </div>
                        </div>
                        <div class="chat-bubble-body">
                          <div v-if="item.type === 0">
                            <p v-html="item.msg.replace(/\n/g, '<br>')"></p>
                          </div>
                          <div v-else-if="item.type === 1">
                            <p v-html="sanitizeMsg(item.msg.replace(/\n/g, '<br>'))"></p>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <!--历史聊天数据渲染 end-->


                <!--停留在页面中聊天内容渲染 start-->
                <div
                  class="chat-item"
                  v-for="item in chatData.list"
                  :key="item.id"
                >
                  <div
                    :class="[
                      'row',
                      item.type === 1
                        ? 'align-items-end justify-content-end'
                        : '',
                    ]"
                  >
                    <div class="col col-lg-9">
                      <div
                        class="chat-bubble"
                        :class="item.type === 1 ? 'chat-bubble-me' : ''"
                      >
                        <div class="chat-bubble-title">
                          <div v-if="item.type === 0">
                            <div class="row">
                              <div class="col chat-bubble-author">机器人</div>
                              <div class="col-auto chat-bubble-date">
                                {{ formatDate(item.created_at) }}
                              </div>
                            </div>
                          </div>
                          <div v-else-if="item.type === 1">
                            <div class="row">
                              <div class="col chat-bubble-date">
                                {{ formatDate(item.created_at) }}
                              </div>
                              <div class="col-auto chat-bubble-author">
                                自己
                              </div>
                            </div>
                          </div>
                        </div>
                        <div class="chat-bubble-body">
                          <div v-if="item.type === 0">
                            <p v-html="item.msg.replace(/\n/g, '<br>')"></p>
                          </div>
                          <div v-else-if="item.type === 1">
                            <p v-html="sanitizeMsg(item.msg.replace(/\n/g, '<br>'))"></p>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <!--停留在页面中聊天内容渲染 end-->

              </div>
            </div>
          </div>
          <div class="card-footer mt-3">
            <div class="input-group input-group-flat">
              <textarea
                id="sendMsgBox"
                ref="sendMsgBox"
                type="text"
                class="form-control"
                autocomplete="off"
                placeholder="Type message"
              ></textarea>
              <span class="input-group-text">
                <a
                  id="sendMsg"
                  href="#"
                  class="link-secondary"
                  data-bs-toggle="tooltip"
                  aria-label="Clear search"
                  data-bs-original-title="Clear search"
                  @click.prevent="sendMessage"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-send"
                  >
                    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                    <path d="M10 14l11 -11" />
                    <path
                      d="M21 3l-6.5 18a.55 .55 0 0 1 -1 0l-3.5 -7l-7 -3.5a.55 .55 0 0 1 0 -1l18 -6.5"
                    />
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
import { ref, reactive } from "vue";

export default {
  name: "ChatPage",
  data() {
    return {
      data: {
        list: [],
        index: 0,
        canFetchMore: true,
      },
    };
  },
  mounted() {
    this.fetchChatRecords(0);
    this.$nextTick(() => {
      const element = document.querySelector(".card-body");
      element.addEventListener("scroll", this.handleScroll);
    });
  },
  methods: {
    sanitizeMsg(msg) {
      // 使用 DOMParser 来解析 HTML 字符串并过滤潜在的恶意内容
      const doc = new DOMParser().parseFromString(msg, 'text/html');
      return doc.body.textContent;
    },
    formatDate(date) {
      return new Date(date).toLocaleString("en-US", {
        month: "2-digit",
        day: "2-digit",
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
      });
    },
    async fetchChatRecords(up) {
      try {
        this.data.canFetchMore = false;
        console.log('111111111111');
        console.log(process.env)
        const response = await fetch(
          `${process.env.VUE_APP_DOMAIN}/api/chat/getChatRecords?start=${this.data.index}`,
          {
            credentials: "include",
          }
        );
        if (response.status === 401) {
          console.log("尚未登陆....");
          this.$router.push("/login");
          return;
        }
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        const res = await response.json();
        this.data.list.push(...res.data.list);
        this.data.index = res.data.index;
        if (res.data.list.length < 20) {
          this.data.canFetchMore = false;
        } else {
          this.data.canFetchMore = true;
        }

        console.log(up);
        // setTimeout(() => {
        //   const element = document.querySelector(".card-body");
        //   element.scrollTop = up == 1 ? 0 : element.scrollHeight;
        // }, 1000);
      } catch (error) {
        console.error("Error fetching chat records:", error);
        // Handle error here, e.g., show a message to the user
      }
    },
    handleScroll(event) {
      const element = event.target;
      if (element.scrollTop <= 50 && this.data.canFetchMore) {
        this.fetchChatRecords(1);
      }
    },
  },
  setup() {
    const sendMsgBox = ref(null);
    const chatData = reactive({
      list: [],
      index: 0,
      canFetchMore: true,
    });

    const sendMessage = async () => {
      const message = sendMsgBox.value.value.trim();

      console.log("sendMsg:----" + message);
      if (message === "") {
        sendMsgBox.value.value = "";
        return;
      }

      chatData.list.push({
        id: Math.floor(Math.random() * 10000),
        type: 1,
        msg: message,
        created_at: Date.now(),
      });

      try {
        const response = await fetch(
          process.env.VUE_APP_DOMAIN+"/api/chat/chat",
          {
            credentials: "include",
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({ msg: message }),
          }
        );

        if (response.ok) {
          const res = await response.json();
          if (res.code === 200) {
            if (res.data.msg) {
              chatData.list.push({
                id: Math.floor(Math.random() * 10000),
                type: 0,
                msg: res.data.msg,
                created_at: Date.now(),
              });
            }

            setTimeout(() => {
              const element = document.querySelector(".card-body");
              element.scrollTop = element.scrollHeight;
            }, 1000);
          }
        }
      } catch (error) {
        console.error("Error sending message:", error);
      }
      sendMsgBox.value.value = "";
    };

    return {
      sendMsgBox,
      sendMessage,
      chatData,
    };
  },
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
