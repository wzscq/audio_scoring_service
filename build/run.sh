docker run -d --name audioscoringservice -p8004:80 -v /root/crvframe/appfile:/services/audioscoringservice/appfile -v /root/audioscoringservice/conf:/services/audioscoringservice/conf  wangzhsh/audioscoringservice:0.0.1