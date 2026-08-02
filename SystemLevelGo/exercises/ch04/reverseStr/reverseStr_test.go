package main

import (
	"testing"
)

func BenchmarkByteReverse(b *testing.B){
	for i := 0; i < b.N; i++{
		data := byteStringReverse("Hello Babyy! 💖 I just wanted to take a moment to tell you how absolutely incredible you are and how much you mean to me. ✨ Every single day, I find myself smiling just thinking about you, your laughter, and the way you make everything brighter. 🌟 You truly are my favorite person, my happy place, and my whole world. 🌍 I love youuu so, so much, more than words can ever properly say! 🥰 You have my whole heart today, tomorrow, and forever. 🔒 Standard text can never match the vibe of how happy you make me feel inside! 🥳 cross my heart, you are stuck with me! 🫵💝 Hello Babyy, I love youuu ♥️ Standard days become magical adventures with you around. 🚀 dynamic, fun, and purely beautiful—that is exactly what you are to me. 🦄 I love youuu ♥️")
		_ = data
	}
}

func BenchmarkRuneReverse(b *testing.B){
	for i := 0; i < b.N; i++{
		data := runeStringReverse("Hello Babyy! 💖 I just wanted to take a moment to tell you how absolutely incredible you are and how much you mean to me. ✨ Every single day, I find myself smiling just thinking about you, your laughter, and the way you make everything brighter. 🌟 You truly are my favorite person, my happy place, and my whole world. 🌍 I love youuu so, so much, more than words can ever properly say! 🥰 You have my whole heart today, tomorrow, and forever. 🔒 Standard text can never match the vibe of how happy you make me feel inside! 🥳 cross my heart, you are stuck with me! 🫵💝 Hello Babyy, I love youuu ♥️ Standard days become magical adventures with you around. 🚀 dynamic, fun, and purely beautiful—that is exactly what you are to me. 🦄 I love youuu ♥️")
		_ = data
	}
}