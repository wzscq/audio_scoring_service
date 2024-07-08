docker run -d --name audioscoringservice -p8004:80 -v /root/crvframe/appfile:/services/audioscoringservice/appfile -v /root/audioscoringservice/conf:/services/audioscoringservice/conf  wangzhsh/audioscoringservice:0.0.1


 GetAudioFile: appfile/audio_test/audio_test_file//original_audio_row36_*
2024/07/02 23:35:16 testrec.go:98: GetAudioFile: appfile/audio_test/audio_test_file//caller_audio_row36_*
2024/07/02 23:35:16 testrec.go:98: GetAudioFile: appfile/audio_test/audio_test_file//called_audio_row36_*

https://github.com/ludlows/PESQ