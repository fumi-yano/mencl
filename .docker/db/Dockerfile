FROM mysql:5.7
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN touch /var/log/mysql/mysqld.log
