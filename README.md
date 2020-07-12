<h1><a id="Employeeledger_0"></a>Employeeledger</h1>
<p>Employeeledger là một web application được xây dựng bằng Go SDK để triển khai một hệ thống mạng Hyper Ledger Fabric. Web app cung cấp một nền tảng các bản ghi người dùng, và một số các chức năng như thêm, sửa xóa người dùng với hai phân quyền chính là người sử dụng (user) và quản trị viên (admin)</p>



<h4><a id="Installation_6"></a>Installation</h4>
<p>Các ứng dụng yêu cầu: <a href="https://www.docker.com/">Docker</a> &amp; <a href="https://golang.org/">Go</a>.</p>
<h3><a id="Docker_10"></a>Docker</h3>
<pre><code class="language-sh">$ sudo apt install docker.io
$ sudo apt install docker-compose
</code></pre>
<h2><a id="Go_15"></a>Go</h2>
<h4><a id="Installation_16"></a>Installation</h4>
<pre><code class="language-sh">$ sudo apt-get update
$ sudo apt-get install golang-go
</code></pre>
<h4><a id="Set_your_Go_path_as_environmental_variable_21"></a>Cài đặt biến môi trường cho Go</h4>
<h6><a id="add_these_following_variable_into_the_profile_22"></a>Thêm các biến môi trường:</h6>
<pre><code class="language-sh">$ <span class="hljs-built_in">export</span> GOPATH=<span class="hljs-variable">$HOME</span>/go
$ <span class="hljs-built_in">export</span> PATH=<span class="hljs-variable">$PATH</span>:/usr/<span class="hljs-built_in">local</span>/go/bin:<span class="hljs-variable">$GOPATH</span>/bin
</code></pre>
<h6><a id="then_27"></a>sau đó</h6>
<pre><code class="language-sh">$ <span class="hljs-built_in">source</span> ~/.profile
$ go version
$ go version go1.<span class="hljs-number">11</span> linux/amd64
</code></pre>



<h3><a id="Build_Your_Network_34"></a>Các bước chạy mạng:</h3>
<p>Mạng Hyperledger Fabric được xây dựng với 1 org và 2 peer</p>


<h4><a id="Run_the_application_90"></a>Chạy ứng dụng</h4>
<ul>
<li>Trước hết cần chuyển thư mục Employeeledger vào đường dẫn <code>home/go/src/github.com</code></li>
<li>Trong thư mục Employeeledger đã có một file "make",mở terminal và gõ "make" và đợi trong giây lát để ứng dụng được build</li>
<li>Khi có thông báo thành công, ứng dụng đã được chạy trên cổng 8000 của localhost</li>
<li>Bây giờ, mở trình duyệt lên và gõ <a href="http://localhost:8000">http://localhost:8000</a></li>
</ul>

<h4>Một số ảnh chụp màn hình</h4>

<h5>Đăng nhập</h5>

<img src="media/export_user_login.png">

<h5>Bảng điều khiển</h5>

<img src="media/dashboard.png">

<h5>Đăng kí</h5>

<img src="media/register.png">

<h5>Đổi mật khẩu</h5>

<img src="media/app_change_password.png">

<h5>Cập nhật thông cá nhân</h5>

<img src="media/update_details.png">

<h2>Video hướng dẫn sử dụng</h2>

<video width="800" height="450" src="https://www.youtube.com/embed/V2hlQkVJZhE" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></video>

