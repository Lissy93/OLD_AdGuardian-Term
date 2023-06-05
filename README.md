<h1 align="center">[OLD] AdGuardian-Term</h1>
<p align="center">
	<i>Terminal-based, real-time traffic monitoring and statistics for your AdGuard Home instance</i><br>
</p>

## About

<img align="right" width="460" src="https://i.ibb.co/BykbdjF/Screenshot-from-2023-06-03-22-57-07.png" >

> **Note**: This app has now been replaced<br>See **[Lissy93/AdGuardian-Term](https://github.com/Lissy93/AdGuardian-Term)** instead!

My purpose for uploading this (failed) project is to show that not everything goes according to plan, and that's okay - it's just part of the learning process!
I've written more about my learnings in the [Background & Lessons Learned](#background--lessons-learned) section below.

Although now re-written, this app is still fully functional, and pretty much feature complete.
If you wish to continue, see the [usage](#usage) instructions for development and deployment docs.

<br><br>

---

## Background & Lessons Learned

Initially, I chose Go, because it's quite performant, has a neat concurrency model, is reasonably quick to write, and (most importantly) compiles into a tidy single binary. I'd also come across the gizak/termui package, which seemed to be the perfect choice for plotting the charts.

However, I found Go's memory safety to be less than ideal. The compiler was irritating and managing threads manually was difficult. I also came across an issue which, for the life of me, I could not resolve - when termui updated the screen, random characters from the previous view would remain, making the charts hard for the user to read.

After spending nearly an entire weekend building out this project, I arrived at the tough realization that I had probably chosen the wrong tech stack. I was initially reluctant to scrap everything I had written so far, but it was the best way forward.

Instead, I picked up Rust - a language I'd been curious about but hadn't found the right opportunity to dive into. Rust, much like Go, is fast and efficient, but with one key difference - it emphasizes safety and manages memory exceptionally well. Rust also compiles into a single binary and has a robust concurrency model. In short, it seemed to address many of the issues I had encountered with Go.

Rebuilding my app in Rust wasn't a walk in the park, though. It has a steep learning curve and some unique concepts, like ownership and borrowing, which took a while to wrap my head around. Nevertheless, the more I delved into it, the more I found myself appreciating its design. In particular, its ability to prevent common programming errors, like null pointer dereferencing and data races, was a huge win for me.

**As developers, we need to remind ourselves that it's okay to take a step back, reassess our decisions, and change our course if necessary. After all, in the long run, it's not about how quickly we can finish a project but about how well we can adapt, learn, and grow along the way. That's how we build not just better software, but also become better developers.**

The new project is uploaded here:

[![Lissy93/AdGuardian-Term - GitHub](https://gh-card.dev/repos/Lissy93/AdGuardian-Term.svg?fullname=)](https://github.com/Lissy93/AdGuardian-Term)


---

## Usage

### Prerequisites
If you haven't already done so, you'll need to [install Go Lang](https://go.dev/doc/install).<br>
Then clone the repo `git clone https://github.com/Lissy93/OLD_AdGuardian-Term.git && cd OLD_AdGuardian-Term`

### Developing
Run `go run .` to start the app

### Deploying
Run `go build` to generate the executable for your system.<br>
You can then launch that with `./adguardian-term`

### Docker
There's also a Dockerfile, although the image is not published to any registry.<br>
This can be built with: `docker build -t adguardian .`<br>
And then run, with: `docker run adguardian`

### Configuring
The app requires details of your AdGuard Home instance (endpoint and credentials).
I didn't get round to making these parameters into environmental variables.
So currently you'll need to set them in [`values.go`](https://github.com/Lissy93/OLD_AdGuardian-Term/blob/main/values/values.go)

---

## License

> _**[Lissy93/OLD_AdGuardian](https://github.com/Lissy93/OLD_AdGuardian)** is licensed under [MIT](https://github.com/Lissy93/OLD_AdGuardian/blob/master/LICENSE) © [Alicia Sykes](https://aliciasykes.com) 2022._<br>
> <sup align="right">For information, see <a href="https://tldrlegal.com/license/mit-license">TLDR Legal > MIT</a></sup>

<details>
<summary>Expand License</summary>

```
The MIT License (MIT)
Copyright (c) Alicia Sykes <alicia@omg.com> 

Permission is hereby granted, free of charge, to any person obtaining a copy 
of this software and associated documentation files (the "Software"), to deal 
in the Software without restriction, including without limitation the rights 
to use, copy, modify, merge, publish, distribute, sub-license, and/or sell 
copies of the Software, and to permit persons to whom the Software is furnished 
to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included install 
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANT ABILITY, FITNESS FOR A
PARTICULAR PURPOSE AND NON INFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```

</details>

---

<!-- License + Copyright -->
<p  align="center">
  <i>© <a href="https://aliciasykes.com">Alicia Sykes</a> 2023</i><br>
  <i>Licensed under <a href="https://gist.github.com/Lissy93/143d2ee01ccc5c052a17">MIT</a></i><br>
  <a href="https://github.com/lissy93"><img src="https://i.ibb.co/4KtpYxb/octocat-clean-mini.png" /></a><br>
  <sup>Thanks for visiting :)</sup>
</p>

<!-- Dinosaur -->
<!-- 
                        . - ~ ~ ~ - .
      ..     _      .-~               ~-.
     //|     \ `..~                      `.
    || |      }  }              /       \  \
(\   \\ \~^..'                 |         }  \
 \`.-~  o      /       }       |        /    \
 (__          |       /        |       /      `.
  `- - ~ ~ -._|      /_ - ~ ~ ^|      /- _      `.
              |     /          |     /     ~-.     ~- _
              |_____|          |_____|         ~ - . _ _~_-_
-->

