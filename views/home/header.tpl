<!DOCTYPE html>
<html lang="utf-8">
<head>
    <meta charset="utf-8">
    <link rel="stylesheet" type="text/css" href="/resources/home/css/style.css">
	<link rel="stylesheet" type="text/css" href="/resources/home/css/wallpaper.css">
    <link rel="stylesheet" href="/resources/layui/css/layui.css?v=5">
    <title>技匠精神</title>
</head>
<body style="background-color:#e0e3ed;">
<style>
	.ikz-nav-item{
		display: flex;
		align-items: center;
	}
	.ikz-nav-item-detail{
		display: flex;
	}
	.ikz-nav-item a{
		margin-right: 16px;
	}
	.ikz-nav-item i{
		font-size: 22px;
		font-weight: 500;
	}
</style>

		<!-- 头部 开始 -->
		<div class="layui-header header trans_3">
			<div class="main index_main">
				<a class="logo" href="/">

				</a>
				<ul class="layui-nav" lay-filter="top_nav">
					<li class="layui-nav-item home">
						<a href="/">首页</a>
					</li>
					{{range .class}}
					<li class="layui-nav-item">
						<a href="?class_id={{.ClassId}}">{{.Name}}</a>
						<dl class="layui-nav-child">
							<!-- 二级菜单 -->
							{{range .Children}}
							<dd>
								<a href="?child_id={{.ClassId}}">{{.Name}}</a>
							</dd>
							{{end}}
						</dl>
					</li>
					{{end}}
					<li class="layui-nav-item">
                    						<a href="/">登录</a>
                    </li>
				</ul>



			</div>
		</div>

        {{/*<div class="header_back"></div>*/}}
        <!-- 头部 结束 -->
        <!-- 左边导航 开始 -->
        <div class="layui-side layui-bg-black left_nav trans_2">
            <div class="layui-side-scroll">
                <ul class="layui-nav layui-nav-tree" lay-filter="left_nav">
                    <li class="layui-nav-item home">
                        <a href="/">首页</a>
                    </li>
                    {{range .class}}
                    <li class="layui-nav-item">
                        <a href="javascript:void();">{{.Name}}</a>
                        <dl class="layui-nav-child">
                            <!-- 二级菜单 -->
                            {{range .Children}}
                                <dd>
                                    <a href="{{.Link}}">{{.Name}}</a>
                                </dd>
                            {{end}}
                        </dl>
                    </li>
                    {{end}}
                </ul>
            </div>
        </div>
        <div class="left_nav_mask"></div>
        <div class="left_nav_btn"><i class="layui-icon">&#xe602;</i></div>
        <!-- 左边导航 结束 -->