<template>
  <el-tabs type="border-card">
    <el-tab-pane label="医院简介">
      <el-card shadow="hover" style="width: 100%; height: 620px;">
        <div slot="header" class="clearfix">
          <span>医院介绍</span>
          <el-divider></el-divider>
        </div>

        <div style="display: flex; align-items: center;">
          <img :src="yourImageUrl" alt="描述性文字" style="width: 500px; height: 400px; margin-right: 20px;" />
          <div style="flex-grow: 1;">
            <p>
              江软大学医学院附属第二医院，作为江软大学医学院的重要组成部分，自建院以来，一直秉承“厚德、博学、求实、创新”的院训，致力于成为医学教育和医疗服务的典范。
              医院拥有一支由资深教授和临床专家组成的医疗团队，专注于提供高质量的医疗服务。医院在心血管病学、神经外科学、肿瘤学和急诊医学等领域具有显著的专业特色。
              作为江软大学医学院的临床教学基地，医院在医学科研方面成绩斐然，拥有多个国家级和省级重点实验室，为医学生和科研人员提供了优越的科研环境。
              医院积极开展国际合作与交流，与世界各地的医学院校建立了合作关系，定期举办国际学术会议，促进了医学知识的交流与传播。
              医院积极参与社会服务活动，通过社区健康讲座、公益义诊等形式，普及健康知识，提升公众健康意识。
              医院坚持以患者为中心，倡导全人医疗和人文关怀，努力为患者提供温馨、舒适的就医环境和专业、高效的医疗服务。
            </p>
          </div>
        </div>

      </el-card>
    </el-tab-pane>

    <el-tab-pane label="医院环境">

      <el-card shadow="hover" style="width: 100%; height: 620px;">
        <div slot="header" class="clearfix">
          <span>医院环境</span>
          <el-divider></el-divider>
        </div>
        <!-- 轮播图组件 -->
        <div class="image-scroll-container">
          <div class="image-list" ref="imageList">
            <el-image v-for="url in urls" :key="url" :src="url" lazy class="image-item"></el-image>
          </div>
        </div>
        <p class="hospital-intro">
          江软大学医学院附属第二医院，以其卓越的医疗环境和先进的设施，为患者提供一个温馨、舒适且专业的治疗和康复空间。医院环境设计充分考虑了患者的生理和心理需求，力求营造一个安全、安静的治疗氛围。
          <span class="hall">宽敞明亮的大厅：</span> 医院的接待大厅宽敞明亮，自然光线充足，为患者提供了一个开阔的视野和舒适的等候体验。大厅内设有清晰的指示牌和导诊服务，方便患者快速找到目的地。
          <span class="equipment">先进的医疗设施：</span> 医院配备了国际先进的医疗设备和仪器，包括最新的影像诊断设备、手术机器人和自动化药房系统。
        </p>
      </el-card>
    </el-tab-pane>


    <el-tab-pane label="医院资讯">

      <el-card shadow="hover" style="width: 100%; height: 100%;">
        <div slot="header" class="clearfix">
          <span>公告框</span>
          <el-divider></el-divider>
        </div>

        <div slot="header" class="fix">
          <span>公告,是指政府、团体对重大事件当众正式公布或者公开宣告,宣布。国务院2012年4月16日发布、2012年7月1日起施行的《党政机关公文处理工作条例》,对公告的使用表述为:“适用于向国内外宣布重要事项或者法定事项”。其中包含两方面的内容:一是向国内外宣布重要事项，公布依据政策、法令采取的重大行动等;二是向国内外宣布法定事项，公布依据法律规定告知国内外的有关重要规定和重大行动等.</span>
         
        </div>

        <el-row>
          <el-col :span="6" v-for="(info, index) in infos" :key="info.id" :offset="index > 0 ? 2 : 0">
            <el-card :body-style="{ padding: '0px' }">
              <!-- 确保图片地址是有效的 -->
              <img :src="info.imageUrl" alt="资讯图片" class="image">
              <div style="padding: 2px;">
                <span>{{ info.title }}</span>
                <div class="bottom clearfix">

                  <div class="stats">
                    <span class="likes">
                      <i class="el-icon-star-off"></i> {{ info.likes }}
                    </span>
                    <span class="views">
                      <i class="el-icon-view"></i> {{ info.views }}
                    </span>
                  </div>

                  <time class="time">{{ currentDate }}</time>
                </div>
              </div>
            </el-card>
          </el-col>
        </el-row>

      </el-card>

    </el-tab-pane>



  </el-tabs>
</template>

<script>
import yourImageUrl from '@/assets/imgs/shouye.jpg';
import image1 from '@/assets/imgs/hosp1.jpg';
import image2 from '@/assets/imgs/hosp2.jpg';
import image3 from '@/assets/imgs/hosp3.jpg';
import image4 from '@/assets/imgs/hosp4.jpg';
import image5 from '@/assets/imgs/hosp5.jpg';
import image6 from '@/assets/imgs/hosp6.jpg';
import { ref } from 'vue';
// import imageUrl1 from '@/assets/imgs/doc1.jpg';
// import imageUrl2 from '@/assets/imgs/doc2.jpg'

export default {
 
  data() {
    return {
      yourImageUrl,

      urls: [
        image1,
        image2,
        image3,
        image4,
        image5,
        image6,
      ],

      // imageUrl1,

      currentDate: new Date(),

         // 资讯列表
      infos: [
        {
          id: 1, 
          imageUrl: '/image/1.png',        
          title: '威海第一针诺西那生钠在威海市立医院完成注射',
          likes: 150,
          views: 1200
        },
        {
          id: 2,
          imageUrl:'/image/2.png',
          title: '江软大学医学院附属医院成功实施了东北三省首例经皮二尖瓣手术。采用了国际上最先进的微创技术，通过皮肤表层的小切口，利用导管将人工瓣膜送至心脏，对病变的二尖瓣进行修复或替换',
          likes: 200,
          views: 1500
        },
        {
          id: 3,
          imageUrl: '/image/3.png',
          title: '江软附属医院还将加强与国内外知名医疗机构的交流与合作，引进更多先进的医疗技术和理念，不断提升医疗服务水平，为患者创造更多的生命奇迹。',
          likes: 200,
          views: 1500
        },
   
      ]
    };
  },


    };


</script>


<style>
 p {
   font-size: 16px;
   /* 设置字体大小为16像素 */
   font-family: Arial, sans-serif;
   /* 设置字体类型为Arial，如果没有Arial则使用无衬线字体 */
   font-weight: normal;
   /* 设置字体粗细为正常 */
   color: #333;
   /* 设置文字颜色为深灰色 */
   line-height: 1.6;
   /* 设置行高为字体大小的1.6倍，提高可读性 */
 }

 /* 可以添加一些自定义样式 */
 .el-carousel__item img {
   width: 100%;
   height: 100%;
   object-fit: cover;
   /* 确保图片覆盖整个轮播项区域，不会变形 */
 }

 /* 环境图 */
 .image-scroll-container {
   overflow-x: auto;
   /* 允许水平滚动 */
   white-space: nowrap;
   /* 防止元素换行 */
   box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
   margin: 20px;
   padding: 10px;
 }

 .image-list {
   display: inline-block;
   /* 使图片列表在一行内显示 */
   padding: 0;
 }

 .image-item {
   margin-right: 20px;
   /* 为图片添加更大的间距 */
   display: inline-block;
   /* 使图片在一行内显示 */
   width: 350px;
   /* 图片宽度，根据需要调整 */
   height: 200px;
   /* 图片高度，根据需要调整 */
   transition: transform 0.5s ease;
   /* 平滑缩放效果 */
   cursor: pointer;
   /* 鼠标悬停时显示指针 */
 }

 .image-item:hover {
   transform: scale(1.05);
   /* 鼠标悬停时图片放大 */
 }



 p.hospital-intro {
   font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
   /* 优雅的无衬线字体 */
   font-size: 18px;
   /* 适中的字体大小 */
   color: #333;
   /* 深色文本，易于阅读 */
   line-height: 1.6;
   /* 良好的行间距，提高可读性 */
   letter-spacing: 0.5px;
   /* 字间距微调 */
   margin: 20px 0;
   /* 段落上下边距 */
   text-align: justify;
   /* 文本两端对齐 */
   hyphens: auto;
   /* 允许连字符 */
 }

 p.hospital-intro::before {
   content: "医院环境介绍";
   /* 标题前缀 */
   display: block;
   font-weight: bold;
   margin-bottom: 10px;
   /* 标题与文本的间距 */
 }

 /* 为特定文本添加样式，例如大厅部分 */
 p.hospital-intro span.hall {
   font-size: 20px;
   /* 稍大的字体大小 */
   color: #0066cc;
   /* 吸引注意的颜色 */
   border-left: 3px solid #0066cc;
   /* 左侧边框 */
   padding-left: 10px;
   /* 内边距 */
   margin-top: 15px;
   /* 与上文的间距 */
   display: block;
   /* 作为块级元素显示 */
 }

 /* 为高级设施部分添加样式 */
 p.hospital-intro span.equipment {
   font-style: italic;
   /* 斜体文本 */
   color: #555;
   /* 稍微淡一些的颜色 */
 }


    .time {
       
      font-size: 13px;
      color: #999;
    }
  
    .bottom {
      margin-top: 13px;
      line-height: 12px;
    }
  
    .button {
      padding: 0;
      float: right;
    }
  
    .image {
      width: 100%;
      height: 200px;
      display: block;
    }
  
    .clearfix:before,
    .clearfix:after {
      display: table;
      content: "";
    }
  
    .clearfix:after {
      clear: both
    }
    .stats {
      overflow: auto;
          /* 清除浮动 */
    }
    
    .likes,
    .views {
      float: left;
          /* 默认向左浮动 */
      margin: 0 5px;
          /* 可选：添加一些间距 */
    }
    
    .likes {
      float: left;
          /* 喜欢图标靠左 */
    }
    
    .views {
      float: right;
          /* 查看图标靠右 */
    }

      .fix {
  background-color: #f5f7fa; /* 公告背景颜色 */
  border: 1px solid #53b3b3; /* 边框颜色和宽度 */
  border-radius: 4px; /* 边框圆角 */
  text-align: center;
  padding: 10px 20px; /* 头部内边距调整，增加左右间距 */
  margin-bottom: 20px; /* 与下方内容的距离 */
}

.fix span {
  font-size: 16px; /* 标题字体大小 */
  color: #1a1b1b; /* 字体颜色改为蓝色 */
  line-height: 1.5; /* 行高 */
  font-weight: bold; /* 字体加粗 */
}

.el-divider--horizontal {
  margin: 10px 0; /* 分隔线上下边距 */
  background-color: #dcdfe6; /* 分隔线颜色 */
}

</style>