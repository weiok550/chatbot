<template>
  <div class="container" style="padding: 0">
    <div class="row">
      <div class="col-12">
        <div class="chatbox">
          <!--历史聊天数据渲染 start-->
          <div
            class="chat-item"
            v-for="item in data.list.slice().reverse()"
            :key="item.id"
          >
            <div
              :class="[
                'row',
                item.type === 1 ? 'align-items-end justify-content-end' : '',
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
                        <div class="col-auto chat-bubble-author">自己</div>
                      </div>
                    </div>
                  </div>
                  <div class="chat-bubble-body">
                    <div v-if="item.type === 0">
                      <p v-html="item.msg.replace(/\n/g, '<br>')"></p>
                    </div>
                    <div v-else-if="item.type === 1">
                      <p
                        v-html="sanitizeMsg(item.msg.replace(/\n/g, '<br>'))"
                      ></p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!--历史聊天数据渲染 end-->

          <!--停留在页面中聊天内容渲染 start-->
          <div class="chat-item" v-for="item in chatData.list" :key="item.id">
            <div
              :class="[
                'row',
                item.type === 1 ? 'align-items-end justify-content-end' : '',
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
                        <div class="col-auto chat-bubble-author">自己</div>
                      </div>
                    </div>
                  </div>
                  <div class="chat-bubble-body">
                    <div v-if="item.type === 0">
                      <p v-html="item.msg.replace(/\n/g, '<br>')"></p>
                    </div>
                    <div v-else-if="item.type === 1">
                      <p
                        v-html="sanitizeMsg(item.msg.replace(/\n/g, '<br>'))"
                      ></p>
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
    <div class="card-footer mt-3 sticky-bottom">
      <div class="input-group input-group-flat">
        <textarea
          id="sendMsgBox"
          ref="sendMsgBox"
          type="text"
          class="form-control"
          autocomplete="off"
          placeholder="Type message"
          @keyup.enter="sendMessage"
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
      const element = document.querySelector(".chatbox");
      element.addEventListener("scroll", this.handleScroll);
    });
  },
  methods: {
    sanitizeMsg(msg) {
      // 使用 DOMParser 来解析 HTML 字符串并过滤潜在的恶意内容
      const doc = new DOMParser().parseFromString(msg, "text/html");
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

        if (up == 0) {
          setTimeout(() => {
            const element = document.querySelector(".chatbox");
            element.scrollTop = element.scrollHeight;
          }, 1000);
        }
      } catch (error) {
        console.error("Error fetching chat records:", error);
      }
    },
    handleScroll(event) {
      const element = event.target;
      if (element.scrollTop <= 100 && this.data.canFetchMore) {
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
      setTimeout(() => {
        const element = document.querySelector(".chatbox");
        element.scrollTop = element.scrollHeight;
      }, 500);

      try {
        const response = await fetch(
          process.env.VUE_APP_DOMAIN + "/api/chat/chat",
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
              const element = document.querySelector(".chatbox");
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
body {
  display: block;
  background-color: #000;
}

.chatbox {
  padding: 20px;
  padding-bottom: 100px;
  height: 100vh;
  background-color: #fff;
  color: #6c7a91;
  font-family: "Quicksand", sans-serif;
  font-size: 14px;
  overflow-y: auto;
  max-height: 100vh; /* 或者你需要的高度 */
}

.chat-item {
  margin: 20px auto;
}

.chat-bubble {
  border-radius: 15px;
  background-color: #eaf0f8;
  padding: 10px;
  margin-bottom: 10px;
}

.chat-bubble-author {
  font-weight: 600;
  text-align: left;
}

.chat-bubble-date {
  text-align: right;
}

.chat-bubble-body {
  margin-top: 10px;
}

.sticky-bottom {
  position: sticky;
  bottom: 0;
}
</style>
