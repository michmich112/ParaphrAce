from parrot import Parrot
import torch
import warnings
warnings.filterwarnings("ignore")


# for reproducibility
def random_state(seed):
  torch.manual_seed(seed)
  if torch.cuda.is_available():
    torch.cuda.manual_seed_all(seed)

random_state(1234)

parrot = Parrot(model_tag="prithivida/parrot_paraphraser_on_T5", use_gpu=False)

while 1:
  print("-"*100)
  phrase = input("Input_phrase: ")
  print("-"*100)
  para_phrases = parrot.augment(
        input_phrase=phrase,
        use_gpu=False,
        diversity_ranker="levenshtein",
        do_diverse=False, 
        max_return_phrases = 10, 
        max_length=32, 
        adequacy_threshold = 0.99, 
        fluency_threshold = 0.80)
  for para_phrase in para_phrases:
   print(para_phrase)
