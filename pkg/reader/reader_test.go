package reader

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testRss = `
<?xml version="1.0"?>
<rss version="2.0">
   <channel>
      <title>Liftoff News</title>
      <link>http://liftoff.msfc.nasa.gov/</link>
      <description>Liftoff to Space Exploration.</description>
      <language>en-us</language>
      <pubDate>Tue, 10 Jun 2003 04:00:00 GMT</pubDate>
      <lastBuildDate>Tue, 10 Jun 2003 09:41:01 GMT</lastBuildDate>
      <docs>http://blogs.law.harvard.edu/tech/rss</docs>
      <generator>Weblog Editor 2.0</generator>
      <managingEditor>editor@example.com</managingEditor>
      <webMaster>webmaster@example.com</webMaster>
      <item>
         <title>Star City</title>
         <link>http://liftoff.msfc.nasa.gov/news/2003/news-starcity.asp</link>
         <description>How do Americans get ready to work with Russians aboard the International Space Station? They take a crash course in culture, language and protocol at Russia's &lt;a href="http://howe.iki.rssi.ru/GCTC/gctc_e.htm"&gt;Star City&lt;/a&gt;.</description>
         <pubDate>Tue, 03 Jun 2003 09:39:21 GMT</pubDate>
         <guid>http://liftoff.msfc.nasa.gov/2003/06/03.html#item573</guid>
         <source url="http://www.tomalak.org/links2.xml">Tomalak's Realm</source>
      </item>
      <item>
         <description>Sky watchers in Europe, Asia, and parts of Alaska and Canada will experience a &lt;a href="http://science.nasa.gov/headlines/y2003/30may_solareclipse.htm"&gt;partial eclipse of the Sun&lt;/a&gt; on Saturday, May 31st.</description>
         <pubDate>Fri, 30 May 2003 11:06:42 GMT</pubDate>
         <guid>http://liftoff.msfc.nasa.gov/2003/05/30.html#item572</guid>
      </item>
      <item>
         <title>The Engine That Does More</title>
         <link>http://liftoff.msfc.nasa.gov/news/2003/news-VASIMR.asp</link>
         <description>Before man travels to Mars, NASA hopes to design new engines that will let us fly through the Solar System more quickly.  The proposed VASIMR engine would do that.</description>
         <pubDate>Tue, 27 May 2003 08:37:32 GMT</pubDate>
         <guid>http://liftoff.msfc.nasa.gov/2003/05/27.html#item571</guid>
      </item>
      <item>
         <title>Astronauts' Dirty Laundry</title>
         <link>http://liftoff.msfc.nasa.gov/news/2003/news-laundry.asp</link>
         <description>Compared to earlier spacecraft, the International Space Station has many luxuries, but laundry facilities are not one of them.  Instead, astronauts have other options.</description>
         <pubDate>Tue, 20 May 2003 08:56:02 GMT</pubDate>
         <guid>http://liftoff.msfc.nasa.gov/2003/05/20.html#item570</guid>
      </item>
   </channel>
</rss>
`

func TestParseRss(t *testing.T) {
	rss, err := parseRss([]byte(testRss))
	require.Nil(t, err)
	require.NotNil(t, rss)

	// Test channel struct
	channel := rss.Channel
	require.NotNil(t, channel)
	assert.Equal(t, "Liftoff News", rss.Channel.Title)
	assert.Equal(t, "Liftoff to Space Exploration.", rss.Channel.Description)
	assert.Equal(t, "http://liftoff.msfc.nasa.gov/", rss.Channel.Link)
	require.Len(t, rss.Channel.Items, 4)

	// Test item struct
	item0 := rss.Channel.Items[0]
	assert.Equal(t, "Star City", item0.Title)
	assert.Equal(t, "How do Americans get ready to work with Russians aboard the International Space Station? "+
		"They take a crash course in culture, language and protocol at Russia's "+
		"<a href=\"http://howe.iki.rssi.ru/GCTC/gctc_e.htm\">Star City</a>.", item0.Description)
	assert.Equal(t, "http://liftoff.msfc.nasa.gov/news/2003/news-starcity.asp", item0.Link)
	utc, err := time.LoadLocation("GMT")
	require.Nil(t, err)
	require.NotNil(t, utc)
	//Tue, 03 Jun 2003 09:39:21 GMT
	assert.True(t, time.Date(2003, time.Month(6), 3, 9, 39, 21, 0, utc).Equal(
		item0.PublishDate.Time))
	assert.Equal(t, "Tomalak's Realm", item0.Source.Title)
	assert.Equal(t, "http://www.tomalak.org/links2.xml", item0.Source.SourceURL)
}