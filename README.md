# Bigly Bot
A  markov chain powered bot to generate new trump tweets. When trained off the provided bullshit.txt, it will sometimes repeat his tweet verbatim: this is just probalistic and related to how unique phrases can be (also due to how new spoken sentences are started).


# Training
Your training file must be new-line delimited strings

ex.
```
I am bigly president
I am a fat unruly child
Yes please I'd like that
How dare you!
Excuse me?
Quickly picked it up it seems
```

I  train my brain with a CSV from [http://www.trumptwitterarchive.com/archive](http://www.trumptwitterarchive.com/archive) and export only the TEXT with no other fields.

# Usage
Will work out of the box with [bullshit.txt](bullshit.txt):

`go run main.go create`

`go run main.go train`

`go run main.go speak`

# Examples

```
00:35 $ go run main.go speak
2019/10/17 00:35:52 Bigly Bot opens his mouth and says...
2019/10/17 00:35:52 Crooked Hillary has the cards but not for this!
```

```
2019/10/17 00:36:06 Bigly Bot opens his mouth and says...
2019/10/17 00:36:06 A mediocre person tells. A good friend: @SarahPalinUSA. More importantly he is the most dishonest media coverage most of her father) in San Diego I have something amazing for the good wishes--you are going back to December 2015. SPYGATE is in Tatters - worst in political history. If the Republicans and Democrats now say using the term Fake News!
```

```
2019/10/17 00:38:14 Bigly Bot opens his mouth and says...
2019/10/17 00:38:14 A big reason is bad for our great Vets Bad for country!
```