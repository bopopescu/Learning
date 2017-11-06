from mrjob.job import MRJob
from mrjob.step import MRStep

class RatingsBreakdown(MRJob):
    def steps(self):
        return [
            MRStep(mapper = self.mapper_get_ratings, reducer = self.reducer_count_rating)
        ]
#take the input line stream, split it by tabs, and return the rating plus the count of 1
    def mapper_get_ratings(self, _, line):
        (userID, movieID, rating, timestamp) = line.split('\t')
        yield rating, 1

#output the key and the sum of all the values for each. Values is an iterator
    def reducer_count_rating(self, key, values):
        yield key, sum(values)

if __name__ == '__main__':
    RatingsBreakdown.run()