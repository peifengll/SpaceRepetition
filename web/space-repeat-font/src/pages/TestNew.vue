<template>
  <vue-ueditor-wrap
      class="new-edit"
      v-model="msg"
      :config="editorConfig"
  ></vue-ueditor-wrap>
  <button>{{msg}} </button>
</template>

<script lang="ts" setup>
const msg = ref('')

//富文本配置
const editorConfig = {
  initialFrameWidth: null,//设置null则宽度自适应
  initialFrameHeight: 320,//区域内容高度
  catchRemoteImageEnable: false, //抓取远程图片是否开启,默认true
  autoHeightEnabled: false, //固定高度且带滚动条
  autoSyncData: false,
  autoFloatEnabled: true,
  serverUrl: "http://192.168.0.194:8080/ueditor/go/controller",
  UEDITOR_HOME_URL: '/UEditor/',
  toolbars: [[
    'fullscreen', 'source', '|', 'undo', 'redo', '|',
    'bold', 'italic', 'underline', 'fontborder', 'strikethrough', 'superscript', 'subscript', 'removeformat', 'formatmatch', 'autotypeset', 'blockquote', 'pasteplain', '|', 'forecolor', 'backcolor', 'insertorderedlist', 'insertunorderedlist', 'selectall', 'cleardoc', '|',
    'rowspacingtop', 'rowspacingbottom', 'lineheight', '|',
    'customstyle', 'paragraph', 'fontfamily', 'fontsize', '|',
    'directionalityltr', 'directionalityrtl', 'indent', '|',
    'justifyleft', 'justifycenter', 'justifyright', 'justifyjustify', '|', 'touppercase', 'tolowercase', '|',
    'link', 'unlink', 'anchor', '|', 'imagenone', 'imageleft', 'imageright', 'imagecenter', '|',
    'simpleupload', 'insertimage', 'emotion', 'scrawl', 'insertvideo', 'music', 'attachment', 'map', 'gmap', 'insertframe', 'insertcode', 'webapp', 'pagebreak', 'template', 'background', '|',
    'horizontal', 'date', 'time', 'spechars', 'snapscreen', 'wordimage', '|',
    'inserttable', 'deletetable', 'insertparagraphbeforetable', 'insertrow', 'deleterow', 'insertcol', 'deletecol', 'mergecells', 'mergeright', 'mergedown', 'splittocells', 'splittorows', 'splittocols', 'charts', '|',
    'print', 'preview', 'searchreplace', 'drafts', 'help'
  ]]
}
</script>
