<template>
  <!-- 外边距: m-2 https://www.tailwindcss.cn/docs/margin -->
  <!-- 内边距： p-4 https://www.tailwindcss.cn/docs/padding  p-0 m-0  -->
  <header class="header py-2 px-4 h-12 w-full flex flex-row justify-between items-center content-center">
    <div class="gap-3 flex flex-row items-center">
      <!-- 返回箭头,点击返回上一页 -->
      <n-tooltip placement="bottom" trigger="hover" v-if="showReturnIcon">
        <template #trigger>
          <i class="icon icon-return" @click="onClickReturnIcon()"></i>
        </template>
        <span>{{ $t("back_button") }}</span>
      </n-tooltip>

      <!-- 服务器设置 -->
      <n-tooltip placement="bottom" trigger="hover">
        <template #trigger>
          <i v-if="!showReturnIcon" class="icon icon-server" role="button" @click="ToAdminPage()"></i>
        </template>
        <span>{{ $t("server_config") }}</span>
      </n-tooltip>

      <!-- 上传按钮，点击进入上传页面 -->
      <n-tooltip placement="bottom" trigger="hover">
        <template #trigger>
          <i class="icon icon-upload" v-if="!showReturnIcon" @click="gotoUploadPage()"></i>
        </template>
        <span>{{ $t("upload_file") }}</span>
      </n-tooltip>

      <!-- 文件重排序 -->
      <n-dropdown v-if="showReSortIcon" trigger="hover" :options="options" @select="onSelect">
        <i class="icon icon-filter"></i>
      </n-dropdown>
    </div>

    <!-- 标题-->
    <!-- 文本颜色： https://www.tailwindcss.cn/docs/text-color -->
    <!-- 文本装饰（下划线）：https://www.tailwindcss.cn/docs/text-decoration -->
    <!-- 文本溢出：https://www.tailwindcss.cn/docs/text-overflow -->
    <!-- 字体粗细:https://www.tailwindcss.cn/docs/font-weight -->
    <div class="p-0 m-0 py-0 flex flex-col justify-center font-semibold content-center truncate">
      <!-- 标题，只显示 -->
      <span class="text-lg" v-if="!setDownLoadLink && (inShelf)">{{ headerTitle }}</span>
      <!-- 标题，可下载压缩包 -->
      <span class="text-lg text-blue-700 text-opacity-100  hover:underline">
        <a v-if="setDownLoadLink && (inShelf)" :href="`api/raw/${bookID}/${encodeURIComponent(headerTitle)}`">{{
          headerTitle
        }}</a>
      </span>
      <!-- 快速跳转 -->
      <QuickJumpBar v-if="!inShelf" class="self-center" :nowBookID="bookID" :readMode="readMode"
        :InfiniteDropdown="InfiniteDropdown"></QuickJumpBar>
    </div>
    <!-- slot，用来插入自定义组件。但是目前没需求 -->
    <!-- <slot></slot> -->

    <!-- 溢出 overflow-x-auton :https://www.tailwindcss.cn/docs/overflow -->
    <div class="gap-3 p-0 h-10 w-33 flex items-center justify-between content-center overflow-x-auton">
      <!-- 扫码阅读，点击可以在屏幕正中显示二维码 -->
      <n-tooltip placement="bottom" trigger="hover">
        <template #trigger>
          <div class="icon icon-qrcode">
            <Qrcode class="w-full h-full opacity-0"></Qrcode>
          </div>
        </template>
        <span>{{ $t("qrcode_hint") }}</span>
      </n-tooltip>

      <!-- 全屏按钮 -->
      <n-tooltip placement="bottom" trigger="hover">
        <template #trigger>
          <div class="p-0 m-0">
            <i class="icon icon-fullscreen" @click="onFullScreen"></i>
          </div>
        </template>
        <span>{{ $t("full_screen_hint") }}</span>
      </n-tooltip>

      <!-- 阅读器设定,点击屏幕中央也可以打开  可自定义方向 -->
      <n-tooltip placement="bottom" trigger="hover">
        <template #trigger>
          <i class="icon icon-setting" v-if="showSettingsIcon" @click="onClickSettingIcon('right')"></i>
        </template>
        <span>{{ $t("ReaderSettings") }}</span>
      </n-tooltip>
    </div>
  </header>
</template>

<script lang="ts">
import axios from "axios";
import screenfull from 'screenfull'
import { h, defineComponent } from 'vue'
import { useCookies } from "vue3-cookies";
import { NIcon, NDropdown, useMessage, NTooltip } from 'naive-ui'
import { ReturnUpBack, SettingsOutline, Grid, List, Filter, Text } from '@vicons/ionicons5'

import Qrcode from "@/components/Qrcode.vue";
import QuickJumpBar from "@/components/QuickJumpBar.vue";

export default defineComponent({
  name: "ComigoHeader",
  props: ['setDownLoadLink', 'headerTitle', 'bookID', 'showReturnIcon', 'showSettingsIcon', 'showReSortIcon', 'readMode', 'inShelf', 'depth', 'InfiniteDropdown'],
  emits: ['drawerActivate', 'onResort'],
  components: {
    NDropdown,//下拉菜单 https://www.naiveui.com/zh-CN/os-theme/components/dropdown
    NIcon,
    //图标,来自 https://www.xicons.org/#/   需要安装（npm i -D @vicons/ionicons5）与导入
    ReturnUpBack,
    Grid,
    List,
    Filter,
    Text,
    SettingsOutline, //图标,来自 https://www.xicons.org/#/   需要安装（npm i -D @vicons/ionicons5）
    Qrcode,//https://github.com/scopewu/qrcode.vue
    QuickJumpBar,
    NTooltip,//https://www.naiveui.com/zh-CN/dark/components/tooltip
  },
  setup() {
    const { cookies } = useCookies();
    // console.log(window.history)
    //警告信息
    const message = useMessage();
    return { cookies, message };
  },
  data() {
    return {
      resort_hint_key: "filename", //书籍的排序方式。可以按照文件名、修改时间、文件大小排序（或反向排序）
      options: [
        {
          label: '卡片模式',
          icon() {
            return h(NIcon, null, {
              default: () => h(Grid)
            })
          },
          key: 'gird'
        },
        {
          label: '列表模式',
          icon() {
            return h(NIcon, null, {
              default: () => h(List)
            })
          },
          key: "list"
        },
        {
          label: '文字模式',
          icon() {
            return h(NIcon, null, {
              default: () => h(Text)
            })
          },
          key: "text"
        },
        {
          type: 'divider',
          key: 'd1'
        },
        {
          label: this.$t("sort_by_filename"),
          key: 'filename'
        },
        {
          label: this.$t("sort_by_modify_time"),
          key: "modify_time"
        },
        {
          label: this.$t("sort_by_filesize"),
          key: 'filesize'
        },
        {
          label: this.$t("sort_by_filename_reverse") + this.$t("sort_reverse"),
          key: 'filename_reverse'
        },
        {
          label: this.$t("sort_by_modify_time_reverse") + this.$t("sort_reverse"),
          key: "modify_time_reverse"
        },
        {
          label: this.$t("sort_by_filesize_reverse") + this.$t("sort_reverse"),
          key: 'filesize_reverse'
        },
      ],
      handleSelect(key: string | number) {
        console.info(String(key))
      },
      showQrcode: false,
    };
  },
  methods: {
    //点击返回图标的时候，后退到上一页或主页
    onClickReturnIcon() {
      axios
        .get("/parent_book_info?id=" + this.bookID)
        .then((response) => {
          this.$router.push("/child_shelf/" + response.data.id);
        }).catch((error) => {
          console.log(error.message);
          this.$router.push('/')
        });
    },
    //根据文件名、修改时间、文件大小等参数重新排序
    onSelect(key: string) {
      // console.info(key);
      this.$emit('onResort', key);
    },
    //进入全屏，由screenfull实现 https://github.com/sindresorhus/screenfull
    //全屏 API： https://developer.mozilla.org/zh-CN/docs/Web/API/Fullscreen_API
    onFullScreen() {
      //如果不允许进入全屏，发提示
      if (!screenfull.isEnabled) {
        this.message.warning(this.$t('not_support_fullscreen'))
        return false
      }
      //切换提示
      if (!screenfull.isFullscreen) {
        this.message.success(this.$t('success_fullScreen'));
      } else {
        this.message.success(this.$t('exit_fullScreen'));
      }
      //切换全屏状态
      screenfull.toggle()
    },

    //点击上传的时候，去上传页
    gotoUploadPage() {
      this.$router.push({
        name: "UploadPage"
      });
    },
    ToAdminPage() {
      location.href = "/admin";
    },

    //点击主页图标的时候，回到主页
    onClickSettingIcon(place: string) {
      this.$emit("drawerActivate", place);
    },

    // 点击 qrcode icon 展示二维码大图
    handleClickQrcode() {
      this.showQrcode = true;
    },
  },
});
</script>

<style lang="less" scoped>
.icon {
  display: block;
  font-size: 28px;
  width: 1em;
  height: 1em;
  background-color: currentColor;
  mask-size: contain;
  mask-repeat: no-repeat;
  mask-position: center;
}

.icon-server {
  mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24' stroke-width='1.5' stroke='currentColor' class='w-6 h-6'%3E%3Cpath stroke-linecap='round' stroke-linejoin='round' d='M5.25 14.25h13.5m-13.5 0a3 3 0 01-3-3m3 3a3 3 0 100 6h13.5a3 3 0 100-6m-16.5-3a3 3 0 013-3h13.5a3 3 0 013 3m-19.5 0a4.5 4.5 0 01.9-2.7L5.737 5.1a3.375 3.375 0 012.7-1.35h7.126c1.062 0 2.062.5 2.7 1.35l2.587 3.45a4.5 4.5 0 01.9 2.7m0 0a3 3 0 01-3 3m0 3h.008v.008h-.008v-.008zm0-6h.008v.008h-.008v-.008zm-3 6h.008v.008h-.008v-.008zm0-6h.008v.008h-.008v-.008z'/%3E%3C/svg%3E");
}

.icon-upload {
  mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' viewBox='0 0 512 512'%3E%3Cpath d='M320 367.79h76c55 0 100-29.21 100-83.6s-53-81.47-96-83.6c-8.89-85.06-71-136.8-144-136.8c-69 0-113.44 45.79-128 91.2c-60 5.7-112 43.88-112 106.4s54 106.4 120 106.4h56' fill='none' stroke='currentColor' stroke-linecap='round' stroke-linejoin='round' stroke-width='32'%3E%3C/path%3E%3Cpath fill='none' stroke='currentColor' stroke-linecap='round' stroke-linejoin='round' stroke-width='32' d='M320 255.79l-64-64l-64 64'%3E%3C/path%3E%3Cpath fill='none' stroke='currentColor' stroke-linecap='round' stroke-linejoin='round' stroke-width='32' d='M256 448.21V207.79'%3E%3C/path%3E%3C/svg%3E");
}

.icon-fullscreen {
  mask-image: url("data:image/svg+xml,%3Csvg class='w-10 static' xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' viewBox='0 0 24 24'%3E%3Cg fill='none' stroke='currentColor' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3E%3Cpath d='M16 4h4v4'%3E%3C/path%3E%3Cpath d='M14 10l6-6'%3E%3C/path%3E%3Cpath d='M8 20H4v-4'%3E%3C/path%3E%3Cpath d='M4 20l6-6'%3E%3C/path%3E%3Cpath d='M16 20h4v-4'%3E%3C/path%3E%3Cpath d='M14 14l6 6'%3E%3C/path%3E%3Cpath d='M8 4H4v4'%3E%3C/path%3E%3Cpath d='M4 4l6 6'%3E%3C/path%3E%3C/g%3E%3C/svg%3E");
}

.icon-qrcode {
  mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' viewBox='0 0 512 512'%3E%3Crect x='336' y='336' width='80' height='80' rx='8' ry='8' fill='currentColor'%3E%3C/rect%3E%3Crect x='272' y='272' width='64' height='64' rx='8' ry='8' fill='currentColor'%3E%3C/rect%3E%3Crect x='416' y='416' width='64' height='64' rx='8' ry='8' fill='currentColor'%3E%3C/rect%3E%3Crect x='432' y='272' width='48' height='48' rx='8' ry='8' fill='currentColor'%3E%3C/rect%3E%3Crect x='272' y='432' width='48' height='48' rx='8' ry='8' fill='currentColor'%3E%3C/rect%3E%3Cpath d='M448 32H304a32 32 0 0 0-32 32v144a32 32 0 0 0 32 32h144a32 32 0 0 0 32-32V64a32 32 0 0 0-32-32zm-32 136a8 8 0 0 1-8 8h-64a8 8 0 0 1-8-8v-64a8 8 0 0 1 8-8h64a8 8 0 0 1 8 8z' fill='currentColor'%3E%3C/path%3E%3Cpath d='M208 32H64a32 32 0 0 0-32 32v144a32 32 0 0 0 32 32h144a32 32 0 0 0 32-32V64a32 32 0 0 0-32-32zm-32 136a8 8 0 0 1-8 8h-64a8 8 0 0 1-8-8v-64a8 8 0 0 1 8-8h64a8 8 0 0 1 8 8z' fill='currentColor'%3E%3C/path%3E%3Cpath d='M208 272H64a32 32 0 0 0-32 32v144a32 32 0 0 0 32 32h144a32 32 0 0 0 32-32V304a32 32 0 0 0-32-32zm-32 136a8 8 0 0 1-8 8h-64a8 8 0 0 1-8-8v-64a8 8 0 0 1 8-8h64a8 8 0 0 1 8 8z' fill='currentColor'%3E%3C/path%3E%3C/svg%3E");
}

.icon-filter {
  mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' viewBox='0 0 512 512'%3E%3Cpath fill='none' stroke='currentColor' stroke-width='32' stroke-miterlimit='10' d='M448 256c0-106-86-192-192-192S64 150 64 256s86 192 192 192s192-86 192-192z'%3E%3C/path%3E%3Cpath fill='none' stroke='currentColor' stroke-width='32' stroke-linecap='round' stroke-linejoin='round' d='M144 208h224'%3E%3C/path%3E%3Cpath fill='none' stroke='currentColor' stroke-width='32' stroke-linecap='round' stroke-linejoin='round' d='M176 272h160'%3E%3C/path%3E%3Cpath fill='none' stroke='currentColor' stroke-width='32' stroke-linecap='round' stroke-linejoin='round' d='M224 336h64'%3E%3C/path%3E%3C/svg%3E");
}

.icon-setting {
  mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' viewBox='0 0 512 512'%3E%3Cpath d='M262.29 192.31a64 64 0 1 0 57.4 57.4a64.13 64.13 0 0 0-57.4-57.4zM416.39 256a154.34 154.34 0 0 1-1.53 20.79l45.21 35.46a10.81 10.81 0 0 1 2.45 13.75l-42.77 74a10.81 10.81 0 0 1-13.14 4.59l-44.9-18.08a16.11 16.11 0 0 0-15.17 1.75A164.48 164.48 0 0 1 325 400.8a15.94 15.94 0 0 0-8.82 12.14l-6.73 47.89a11.08 11.08 0 0 1-10.68 9.17h-85.54a11.11 11.11 0 0 1-10.69-8.87l-6.72-47.82a16.07 16.07 0 0 0-9-12.22a155.3 155.3 0 0 1-21.46-12.57a16 16 0 0 0-15.11-1.71l-44.89 18.07a10.81 10.81 0 0 1-13.14-4.58l-42.77-74a10.8 10.8 0 0 1 2.45-13.75l38.21-30a16.05 16.05 0 0 0 6-14.08c-.36-4.17-.58-8.33-.58-12.5s.21-8.27.58-12.35a16 16 0 0 0-6.07-13.94l-38.19-30A10.81 10.81 0 0 1 49.48 186l42.77-74a10.81 10.81 0 0 1 13.14-4.59l44.9 18.08a16.11 16.11 0 0 0 15.17-1.75A164.48 164.48 0 0 1 187 111.2a15.94 15.94 0 0 0 8.82-12.14l6.73-47.89A11.08 11.08 0 0 1 213.23 42h85.54a11.11 11.11 0 0 1 10.69 8.87l6.72 47.82a16.07 16.07 0 0 0 9 12.22a155.3 155.3 0 0 1 21.46 12.57a16 16 0 0 0 15.11 1.71l44.89-18.07a10.81 10.81 0 0 1 13.14 4.58l42.77 74a10.8 10.8 0 0 1-2.45 13.75l-38.21 30a16.05 16.05 0 0 0-6.05 14.08c.33 4.14.55 8.3.55 12.47z' fill='none' stroke='currentColor' stroke-linecap='round' stroke-linejoin='round' stroke-width='32'%3E%3C/path%3E%3C/svg%3E");
}

.icon-return {
  mask-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' viewBox='0 0 512 512'%3E%3Cpath fill='none' stroke='currentColor' stroke-linecap='round' stroke-linejoin='round' stroke-width='32' d='M112 160l-64 64l64 64'%3E%3C/path%3E%3Cpath d='M64 224h294c58.76 0 106 49.33 106 108v20' fill='none' stroke='currentColor' stroke-linecap='round' stroke-linejoin='round' stroke-width='32'%3E%3C/path%3E%3C/svg%3E");
}
</style>




