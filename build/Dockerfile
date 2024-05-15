FROM alpine:3.15
RUN set -eux &&sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add nginx && apk add --update python3 && apk add py3-pip && pip3 install librosa && pip3 install pesq && apk add ffmpeg
COPY nginx.conf /etc/nginx/nginx.conf
#ADD ./package/web /web
ADD ./package/service/audio_scoring_service /services/audioscoringservice/audio_scoring_service
ADD audioscore.py /services/audioscoringservice/audioscore.py
copy entrypoint.sh /entrypoint.sh
ENTRYPOINT ["sh","entrypoint.sh"]